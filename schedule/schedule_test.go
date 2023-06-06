package schedule

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gopkg.in/go-mixed/framework.v1/contracts/schedule"
	"gopkg.in/go-mixed/framework.v1/testing/mock"
)

func TestApplication(t *testing.T) {
	mockArtisan := mock.Artisan()
	mockArtisan.On("Call", "test --name Laravel argument0 argument1").Return().Times(3)

	immediatelyCall := 0
	delayIfStillRunningCall := 0
	skipIfStillRunningCall := 0

	app := NewSchedule()
	app.Register([]schedule.Event{
		app.Call(func() {
			immediatelyCall++
		}).EveryMinute(),
		app.Call(func() {
			time.Sleep(61 * time.Second)
			delayIfStillRunningCall++
		}).EveryMinute().DelayIfStillRunning(),
		app.Call(func() {
			time.Sleep(61 * time.Second)
			skipIfStillRunningCall++
		}).EveryMinute().SkipIfStillRunning(),
		app.Command("test --name Laravel argument0 argument1").EveryMinute(),
	})

	second, _ := strconv.Atoi(time.Now().Format("05"))
	// Make sure run 3 times
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(120+6+60-second)*time.Second)
	go func(ctx context.Context) {
		app.Run()

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	time.Sleep(time.Duration(120+5+60-second) * time.Second)

	assert.Equal(t, 3, immediatelyCall)
	assert.Equal(t, 2, delayIfStillRunningCall)
	assert.Equal(t, 1, skipIfStillRunningCall)
	mockArtisan.AssertExpectations(t)
}
