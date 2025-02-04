package problem8

func multiplex(chs []<-chan string) []string {
	out := make([]string, 0)
	openChannels := len(chs)

	for openChannels > 0 {
		for i, ch := range chs {
			if ch == nil {
				continue
			}

			select {
			case val, ok := <-ch:
				if ok {
					out = append(out, val)
				} else {
					chs[i] = nil
					openChannels--
				}
			default:
			}
		}
	}
	return out
}
