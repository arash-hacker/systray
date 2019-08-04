package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"systray"
	"time"
	//"github.com/getlantern/systray"
)

func main() {
	onExit := func() {
		fmt.Println("Starting onExit")
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
		fmt.Println("Finished onExit")
	}
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

func onReady() {
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	/* go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
	*/
	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(getIcon("assets/hacker.ico"))
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")

		systray.AddMenuItem("Ignored", "Ignored")

		mUrl := systray.AddMenuItem("Open Lantern.org", "my home")

		mUrlSub := mUrl.AddSubMenuItem("mUrl-Sub1", "mUrl-Sub1")
		/*      */ mUrl.AddSubMenuItem("mUrl-Sub2", "mUrl-Sub2")
		/*      */ mUrl.AddSubMenuItem("mUrl-Sub3", "mUrl-Sub3")

		mUrilSubSub := mUrlSub.AddSubMenuItem("mUrl-Sub1-Sub1", "mUrl-Sub1-Sub1")
		/*          */ mUrlSub.AddSubMenuItem("mUrl-Sub1-Sub2", "mUrl-Sub1-Sub2")
		/*          */ mUrlSub.AddSubMenuItem("mUrl-Sub1-Sub3", "mUrl-Sub1-Sub3")
		/*          */ mUrlSub.AddSubMenuItem("mUrl-Sub1-Sub4", "mUrl-Sub1-Sub4")
		mUrilSubSub.AddSubMenuItem("mUrl-Sub-Sub-Sub", "mUrl-Sub-Sub-Sub")

		mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		mQuit.SetIcon(getIcon("assets/hacker.ico"))

		systray.AddSeparator()
		mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
		shown := true
		for {
			select {

			case <-mUrlSub.ClickedCh:
				fmt.Println("murl")
			case <-mUrilSubSub.ClickedCh:
				fmt.Println("murlSubSub")

			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
				fmt.Println("AAA")

			case <-mUrlSub.ClickedCh:
				mUrlSub.SetTitle("I've Changed")

			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-mToggle.ClickedCh:
				if shown {
					mQuitOrig.Hide()
					mEnabled.Hide()
					shown = false
				} else {
					mQuitOrig.Show()
					mEnabled.Show()
					shown = true
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
	go func() {
		for {
			systray.SetTitle(getTime())
			//systray.SetTooltip("Look at me, I'm a tooltip!")
			time.Sleep(1 * time.Second)
			//systray.AddMenuItem("退出", "Quit the whole app")
		}
	}()

}
func getTime() string {
	t := time.Now()
	hour, min, sec := t.Clock()
	return ItoaTwoDigits(hour) + ":" + ItoaTwoDigits(min) + ":" + ItoaTwoDigits(sec)
}
func ItoaTwoDigits(i int) string {
	b := "0" + strconv.Itoa(i)
	return b[len(b)-2:]
}
func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
