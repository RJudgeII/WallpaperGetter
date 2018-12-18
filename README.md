# WallpaperGetter

This program will get the current "Top of the last 24 hours" posts on the reddit.com/r/wallpapers subreddit. It will look through the posts and find the top that is just a single image (just in case an album of images is the current top), download the image, and set your desktop wallpaper to the image.

## Getting Started

As it stands, this program will only work for Windows desktops. I may add in support for other OSs in the future. To use it, simply clone the repository, run `go run main.go` in the base directory, or build the project with `go build main.go` and run the exe file. I like to build it and add the exe to the `shell:startup` folder so that any time I startup my desktop there will be a new background.