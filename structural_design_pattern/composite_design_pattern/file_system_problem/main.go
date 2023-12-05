package main

func clientCode(system FileSystem) {
	system.ls()
}

func main() {
	reportFile := NewFile("Report.doc")
	picturesFile := NewFile("Picture.jpg")
	files := []FileSystem{reportFile, picturesFile}
	documentFolder := NewFolder("Documents", files)
	clientCode(documentFolder)
}
