package oto

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

var o *Oto

type Oto struct {
	rate   int
	ctx    *oto.Context
	player *oto.Player
}

func new(rate int) (*Oto, error) {
	o := &Oto{}
	var err error

	o.rate = rate
	o.ctx, err = oto.NewContext(rate, 2, 2, 8192)
	if err != nil {
		return nil, err
	}
	o.player = o.ctx.NewPlayer()

	return o, err
}

//Play open an audio file and play its content
func Play(filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		f, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer f.Close()

		d, err := mp3.NewDecoder(f)
		if err != nil {
			return err
		}

		//create oto according to file sample rate.
		//if different sample rate, then re-create oto.
		if o == nil {
			o, err = new(d.SampleRate())
			if err != nil {
				return err
			}
		} else if o.rate != d.SampleRate() {
			o.player.Close()
			o.ctx.Close()

			o, err = new(d.SampleRate())
			if err != nil {
				return err
			}
		}

		if _, err := io.Copy(o.player, d); err != nil {
			return err
		}

		return nil
	} else if os.IsNotExist(err) {
		return err
	} else {
		return errors.New("Schrodinger state: file may or may not exist..")
	}
}

//Close will close context which would close player as well
func Close() {
	if o != nil {
		log.Println("[pkg][oto] closing..")
		o.player.Close()
		o.ctx.Close()
	}
	log.Println("[pkg][oto] done closing")
}
