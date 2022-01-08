package sound

// import (
// 	"bytes"
// 	_ "embed"
// 	"fmt"
// 	"io"
// 	"time"

// 	"github.com/faiface/beep"
// 	"github.com/faiface/beep/speaker"
// 	"github.com/faiface/beep/wav"
// )

// //go:embed scan.wav
// var scanFile []byte

// //go:embed startup.wav
// var startupFile []byte

// func getWavSoundBuffer(source io.Reader) (*beep.Buffer, error) {
// 	streamer, format, err := wav.Decode(source)
// 	if err != nil {
// 		return nil, err
// 	}

// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))

// 	buffer := beep.NewBuffer(format)
// 	buffer.Append(streamer)
// 	streamer.Close()
// 	return buffer, nil
// }
// func PlayScanSound() error {
// 	buffer, err := getWavSoundBuffer(bytes.NewReader(scanFile))
// 	if err != nil {
// 		return fmt.Errorf("error playing scan sound: %v", err)
// 	}
// 	if buffer == nil {
// 		return fmt.Errorf("no sound buffer")
// 	}
// 	shot := buffer.Streamer(0, buffer.Len())
// 	speaker.Play(shot)
// 	return nil
// }
// func PlayStartupSound() error {
// 	buffer, err := getWavSoundBuffer(bytes.NewReader(startupFile))
// 	if err != nil {
// 		return fmt.Errorf("error playing startup sound: %v", err)
// 	}
// 	if buffer == nil {
// 		return fmt.Errorf("no sound buffer")
// 	}
// 	shot := buffer.Streamer(0, buffer.Len())
// 	speaker.Play(shot)
// 	return nil
// }
