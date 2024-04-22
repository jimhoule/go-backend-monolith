package services

import (
	"bufio"
	"fmt"
	"os/exec"
)

type FfmpegTranscoderService struct{}

func (fts *FfmpegTranscoderService) TranscodeDash(dirPath string, fileName string, fileExtension string) error {
	// Creates DASH files in temp dir based on video in same temp dir
	command := exec.Command(
		"ffmpeg",
		// Read input at native frame rate (This is equivalent to setting -readrate 1)
		"-re",
		// Input
		"-i", fmt.Sprintf("%s/%s.%s", dirPath, fileName, fileExtension),
		// Map
		"-map", "0",
		"-map", "0",
		// Codec audio
		"-c:a", "aac",
		// Codec video
		"-c:v", "libx264",
		// Bitrate video
		"-b:v:0", "800k",
		"-b:v:1", "300k",
		// Sampling Rate audio
		"-ar:a:0", "48000",
		"-ar:a:1", "22050",
		// Resolution video
		"-s:v:1", "320x170",
		// Profile video
		"-profile:v:0", "main",
		"-profile:v:1", "baseline",
		// Minimum GOP size
		"-keyint_min", "120",
		// Keyframe interval (A keyframe is inserted at least every -g frames, sometimes sooner)
		"-g", "120",
		// Sets the threshold for the scene change detection
		"-sc_threshold", "0",
		// Sets the maximum number of B frames (1,3,7,15)
		"-bf", "1",
		// Adaptive B-frame placement decision algorithm (Use only on first-pass)
		"-b_strategy", "0",
		// Enables or disables use of SegmentTemplate instead of SegmentList in the manifest (This is enabled by default)
		"-use_template", "1",
		// Enable or disable use of SegmentTimeline within the SegmentTemplate manifest section (This is enabled by default)
		"-use_timeline", "1",
		// Maximum number of segments kept in the manifest, discard the oldest one and when set to 0, all segments are kept (Good for live streaming)
		"-window_size", "0",
		// Enables or disables storing all segments in one file, accessed using byte ranges (This is disabled by default)
		"-single_file", "1",
		// Set of one or more streams accessed as a single subset specified in manifest (.mpd)
		"-adaptation_sets", "id=0,streams=v id=1,streams=a",
		// Forced output file format
		"-f", "dash",
		fmt.Sprintf("%s/%s.mpd", dirPath, fileName),
	)

	// Gets stderr pipe that will be connected to command's standard output when command starts
	// NOTE: For some reason all of ffmpeg outputs are from stderr so we have to use this one
	stderr, err := command.StderrPipe()
	if err != nil {
		return err
	}

	// Starts command
	err = command.Start()
	if err != nil {
		return err
	}

	// Creates scanner
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)

	// Gets command output in real time
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	// Waits for command to exit
	err = command.Wait()
	if err != nil {
		return err
	}

	return nil
}