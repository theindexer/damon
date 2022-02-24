package view

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/hcjulz/damon/component"
	"github.com/hcjulz/damon/models"
	"github.com/hcjulz/damon/styles"
)

func (v *View) Logs(tasks []string, allocID, source string) {
	if len(tasks) > 1 {
		v.components.SelectorModal.Props.Items = tasks
		v.components.SelectorModal.Props.AllocationID = allocID
		v.components.SelectorModal.SetSelectedFunc(func(task string) {
			v.Logs([]string{task}, allocID, source)
		})

		v.url = fmt.Sprintf("allocations/%s/%s/logs", allocID, source)

		v.components.SelectorModal.Render()

		v.Draw()

		v.Layout.Container.SetFocus(v.components.SelectorModal.Modal.Primitive())
		return
	}

	taskName := tasks[0]

	v.Layout.Body.Clear()

	v.Layout.Container.SetInputCapture(v.InputLogs)
	v.components.Commands.Update(component.LogCommands)

	update := func() {
		v.components.LogStream.Props.Data = filterLogs(v.state.Logs, v.state.Filter.Logs)
		v.components.LogStream.Render()

		v.Draw()

		v.Layout.Container.SetFocus(v.components.LogStream.TextView.Primitive())
	}

	v.Watcher.SubscribeToLogs(allocID, taskName, source, update)

	update()

	v.addToHistory(v.state.SelectedNamespace, models.TopicLog, func() {
		update()
	})

	v.Layout.Container.SetInputCapture(v.InputLogs)
}

func filterLogs(logs []byte, filter string) []byte {
	buf := bytes.Buffer{}
	defer buf.Reset()

	if filter != "" {
		rx, _ := regexp.Compile(filter)
		logLines := bytes.Split(logs, []byte("\n"))
		var result []byte
		for _, log := range logLines {
			if rx.Match(log) {
				idx := rx.FindIndex([]byte(log))
				fmt.Fprintf(
					&buf,
					"%s%s%s%s%s%s\n",
					[]byte(styles.ColorLighGreyTag),
					log[:idx[0]],
					[]byte(styles.HighlightSecondaryTag),
					log[idx[0]:idx[1]],
					[]byte(styles.ColorLighGreyTag),
					log[idx[1]:],
				)

				result = append(result, buf.Bytes()...)
				buf.Reset()
			}
		}
		return result
	}

	fmt.Fprintf(&buf, "%s%s", styles.ColorWhiteTag, logs)

	return buf.Bytes()
}
