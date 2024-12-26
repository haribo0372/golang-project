package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("\tTask-1")
	ch := make(chan int)
	go count(ch)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	fmt.Println("\n\tTask-2")
	drawImg1 := openPngAndReturnRGBA("image1.png")

	start := time.Now()
	filter(drawImg1)
	duration := time.Since(start)
	fmt.Println("Время не параллельной обработки:", duration)

	saveRGBAToPng("output_not_parallel", drawImg1)
	fmt.Println("Завершено и сохранено в output_not_parallel.png")

	fmt.Println("\n\tTask-3")
	drawImg2 := openPngAndReturnRGBA("image1.png")

	start = time.Now()
	filterImageParallel(drawImg2)
	duration = time.Since(start)
	fmt.Println("Время параллельной обработки:", duration)

	saveRGBAToPng("output_parallel", drawImg2)
	fmt.Println("Завершено и сохранено в output_parallel.png")

	fmt.Println("\n\tTask-4")
	drawImg3 := openPngAndReturnRGBA("image1.png")

	// Mатрицa размытия (сумма весов = 1)
	kernel := [][]float64{
		{0.0625, 0.125, 0.0625},
		{0.125, 0.25, 0.125},
		{0.0625, 0.125, 0.0625},
	}

	start = time.Now()
	outputImg := applyMatrixFilterImageParallel(drawImg3, kernel)
	duration = time.Since(start)
	fmt.Println("Время параллельной обработки с матричным фильтром:", duration)

	saveRGBAToPng("output_matrix_filter", outputImg)
	fmt.Println("Завершено и сохранено в output_matrix_filter.png")
}

func filter(img draw.Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y).(color.RGBA)

			gray := uint16((uint32(originalColor.R) + uint32(originalColor.G) + uint32(originalColor.B)) / 3)
			img.Set(x, y, color.RGBA{
				R: uint8(gray),
				G: uint8(gray),
				B: uint8(gray),
				A: originalColor.A,
			})
		}
	}
}

func filterParallel(img *image.RGBA, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := img.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		originalColor := img.At(x, y).(color.RGBA)

		gray := uint16((uint32(originalColor.R) + uint32(originalColor.G) + uint32(originalColor.B)) / 3)

		img.Set(x, y, color.RGBA{
			R: uint8(gray),
			G: uint8(gray),
			B: uint8(gray),
			A: originalColor.A,
		})
	}
}

// Обработка изображения параллельно по строкам
func filterImageParallel(img *image.RGBA) {
	bounds := img.Bounds()
	height := bounds.Max.Y - bounds.Min.Y

	var wg sync.WaitGroup
	wg.Add(height)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		go filterParallel(img, y, &wg)
	}

	wg.Wait()
}

func count(ch <-chan int) {
	for num := range ch {
		result := num * num
		fmt.Println(result)
	}
}

func openPngAndReturnRGBA(imageName string) *image.RGBA {
	file, err := os.Open(imageName)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return nil
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Ошибка расшифровки PNG:", err)
		return nil
	}

	drawImg, ok := img.(*image.RGBA)
	if !ok {
		fmt.Println("Ошибка преобразования в *image.RGBA")
		return nil
	}

	return drawImg
}

func saveRGBAToPng(outputFileName string, drawImg *image.RGBA) {
	outputFile, err := os.Create(outputFileName + ".png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка кодирования в PNG:", err)
		return
	}
}

// Применение матричного фильтра к одной строке изображения
func applyMatrixFilterParallel(src, dst *image.RGBA, kernel [][]float64, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := src.Bounds()
	kernelSize := len(kernel)
	kernelOffset := kernelSize / 2

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var rSum, gSum, bSum float64
		for ky := 0; ky < kernelSize; ky++ {
			for kx := 0; kx < kernelSize; kx++ {
				sampleX := clamp(x+kx-kernelOffset, bounds.Min.X, bounds.Max.X-1)
				sampleY := clamp(y+ky-kernelOffset, bounds.Min.Y, bounds.Max.Y-1)
				originalColor := src.At(sampleX, sampleY).(color.RGBA)

				weight := kernel[ky][kx]
				rSum += float64(originalColor.R) * weight
				gSum += float64(originalColor.G) * weight
				bSum += float64(originalColor.B) * weight
			}
		}

		dst.Set(x, y, color.RGBA{
			R: uint8(clampInt(int(rSum), 0, 255)),
			G: uint8(clampInt(int(gSum), 0, 255)),
			B: uint8(clampInt(int(bSum), 0, 255)),
			A: src.At(x, y).(color.RGBA).A,
		})
	}
}

// Применение матричного фильтра параллельно ко всем строкам изображения
func applyMatrixFilterImageParallel(src *image.RGBA, kernel [][]float64) *image.RGBA {
	bounds := src.Bounds()
	height := bounds.Max.Y - bounds.Min.Y

	dst := image.NewRGBA(bounds)
	var wg sync.WaitGroup
	wg.Add(height)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		go applyMatrixFilterParallel(src, dst, kernel, y, &wg)
	}

	wg.Wait()
	return dst
}

// Вспомогательная функция для ограничения значения в пределах min и max
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Вспомогательная функция для ограничения значения (float64 -> int)
func clampInt(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
