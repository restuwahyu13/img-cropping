package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

/*
NOTE: Fungsi ini digunakan untuk memuat gambar dari lokasi dimana file tersebut berada dan mengubah file menjadi sebuah image binary
namaFile: nama file
*/

func parsingGambar(file []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	return img, nil
}

/*
NOTE: Fungsi ini untuk konversi warna 16 bit, ke warna 8 bit standar RGBA dimulai dari 0 - 255, kemudian warna tersebut akan di kompare dengan warna target

c: warna yang akan di konversi
target: warna target
*/
func kompareWarnaTarget(c color.Color, target color.RGBA) bool {
	r, g, b, a := c.RGBA()
	return r/257 == uint32(target.R) &&
		g/257 == uint32(target.G) &&
		b/257 == uint32(target.B) &&
		a/257 == uint32(target.A)
}

/*
NOTE: Fungsi ini untuk mencari kotak hitam yang berada di gambar berdasar sumbu x dan y

gambar: metadata gambar
*/
func cariKotakHitam(gambar image.Image) (*image.Rectangle, error) {
	var (
		simpanMaxArea  int              = 0
		kordinat       image.Rectangle  = gambar.Bounds()
		warnaTarget    color.RGBA       = color.RGBA{0, 0, 0, 255}
		simpanKordinat *image.Rectangle = new(image.Rectangle)
	)

	for y := kordinat.Min.Y; y < kordinat.Max.Y; y++ {
		for x := kordinat.Min.X; x < kordinat.Max.X; x++ {
			if !kompareWarnaTarget(gambar.At(x, y), warnaTarget) {
				continue
			}

			kotak := cariTargetArea(gambar, x, y)
			if area := kotak.Dx() * kotak.Dy(); area > simpanMaxArea {
				simpanMaxArea = area
				simpanKordinat = &kotak
			}
		}
	}

	buatLog(simpanKordinat)

	if simpanMaxArea == 0 {
		return nil, errors.New("kotak hitam tidak ditemukan")
	} else if simpanKordinat.Dx() < 50 || simpanKordinat.Dy() < 50 {
		return nil, errors.New("kotak hitam tidak ditemukan")
	}

	return simpanKordinat, nil
}

/*
NOTE: Fungsi ini digunakan untuk mencari target yang kita inginkan berdasarkan kordinat dari sumbu x dan y

gambar: metadata gambar
x: kordinat sumbu x
y: kordinat sumbu y
*/
func cariTargetArea(gambar image.Image, x, y int) image.Rectangle {
	warna := gambar.At(x, y)

	kanan := x
	for kanan < gambar.Bounds().Max.X && gambar.At(kanan, y) == warna {
		kanan++
	}

	bawah := y
	for bawah < gambar.Bounds().Max.Y && gambar.At(x, bawah) == warna {
		bawah++
	}

	return image.Rect(x, y, kanan, bawah)
}

/*
NOTE: Fungsi ini digunakan untuk membuat gambar baru berdasarkan kordinat yang telah di tentukan

gambar: metadata gambar
rect: kordinat baru yang akan dibuat
*/
func potongGambar(gambar image.Image, rect image.Rectangle) image.Image {
	potong := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			potong.Set(x, y, gambar.At(x, y))
		}
	}

	return potong
}

/*
NOTE: Fungsi ini digunakan untuk menyimpan gambar kedalam directory

namaFile: nama file yang akan dibuat
gambar: metadata gambar
*/
func simpanGambar(namaFile string, gambar image.Image) {
	file, err := os.Create(namaFile)
	if err != nil {
		log.Fatal("Gagal membuat gambar: ", err)
	}
	defer file.Close()

	enc := png.Encoder{}
	enc.CompressionLevel = png.BestCompression

	if err := enc.Encode(file, gambar); err != nil {
		log.Fatal("Gagal menyimpan gambar: ", err)
	}
}

/*
NOTE: Fungsi ini digunakan untuk mengcapture kordinat dari kotak hitam

rect: kordinat baru yang akan dibuat
*/
func buatLog(rect *image.Rectangle) {
	kordinat := fmt.Sprintf("Position X: %d dan Position Y: %d\n", rect.Dx(), rect.Dy())

	file, err := os.OpenFile("image.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Gagal membuat file log: ", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("\n%s", kordinat)); err != nil {
		log.Fatal("Gagal menulis ke log: ", err)
	}
}
