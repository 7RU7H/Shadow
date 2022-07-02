package ninjashell

import ( 
	""
)


//Create a progress bar for the file transfer in Stdout for listener
func createProgressBar(fileSize int) {
	bar := pb.New(fileSize)
	bar.SetMaxWidth(80)
	bar.SetRefreshRate(time.Millisecond * 10)
	bar.Start()
	return bar
}

//Update the progress bar for the file transfer in Stdout for listener
func updateProgressBar(bar *pb.ProgressBar, nBytes int) {
	bar.Increment()
}
