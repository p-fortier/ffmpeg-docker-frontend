package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleCatVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: generateCatVideo")
	//catNumber := strings.TrimPrefix(r.URL.Path, "/cat")
	//number, err := strconv.ParseInt(catNumber, 10, 64)
	//if err != nil {
	//	_, _ = fmt.Fprintf(w, "Error %d", err)
	//	return
	//}

	//Create a temp folder/directory at a full qualified path
	err := os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}

	number := int64(5)
	cats := make([]image.Image, 0)

	for i := int64(0); i < number; i++ {
		res, err := http.Get(serverUrl + "/cat")
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error 1 %d", err)
			return
		}
		cat, err := png.Decode(res.Body)
		if err != nil {
			fmt.Fprintf(w, "error 2 %d", err)
		}

		err = savePicture(cat, i)
		if err != nil {
			fmt.Fprintf(w, "error 3 %d", err)
		}
		cats = append(cats, cat)
		err = res.Body.Close()
		if err != nil {
			return
		}
		println("cats size " + fmt.Sprint(len(cats)))
	}

	//	ffmpeg -r 60 -f image2 -s 1920x1080 -i cat_%d.png -vcodec libx264 -crf 25  -pix_fmt yuv420p output.mp4

	//remove temp dir
	os.ReadDir("temp")
}

func createCatvideo()
func savePicture(cat image.Image, n int64) error {
	f, _ := os.Create("temp/cat_" + fmt.Sprint(n))
	err := png.Encode(f, cat)
	if err != nil {
		return err
	}
	return nil
}

var serverUrl = "http://localhost:4567"

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cat", handleCatVideo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
