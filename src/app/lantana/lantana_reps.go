package lantana

import (
	"context"
	"fmt"
	"sync"

	"github.com/mt3hr/rykv/kyou"
)

type LantanaReps []LantanaRep

func (l LantanaReps) GetAllLantanas(ctx context.Context) ([]*Lantana, error) {
	lantanas := map[string]*Lantana{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan []*Lantana, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			matchLantanas, err := rep.GetAllLantanas(ctx)
			if err != nil {
				errch <- err
				return
			}
			ch <- matchLantanas
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at get all lantanas: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return nil, err
	}

loop:
	for {
		select {
		case k := <-ch:
			if k == nil {
				continue loop
			}
			for _, lantana := range k {
				lantanas[lantana.LantanaID] = lantana
			}
		default:
			break loop
		}
	}

	allLantanas := []*Lantana{}
	for _, lantana := range lantanas {
		if lantana == nil {
			continue
		}
		allLantanas = append(allLantanas, lantana)
	}
	return allLantanas, nil
}

func (l LantanaReps) GetLantana(ctx context.Context, lantanaID string) (*Lantana, error) {
	lantanas := []*Lantana{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan *Lantana, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			matchLantana, err := rep.GetLantana(ctx, lantanaID)
			if err != nil {
				// errch <- err
				return
			}
			ch <- matchLantana
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at get lantana: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return nil, err
	}

loop:
	for {
		select {
		case lantana := <-ch:
			if lantana == nil {
				continue loop
			}
			lantanas = append(lantanas, lantana)
		default:
			break loop
		}
	}

	allLantanas := []*Lantana{}
	for _, lantana := range lantanas {
		if lantana == nil {
			continue
		}
		allLantanas = append(allLantanas, lantana)
	}
	for _, lantana := range allLantanas {
		return lantana, nil
	}
	return nil, fmt.Errorf("not found lantana %s", lantanaID)
}

func (l LantanaReps) AddLantana(ctx context.Context, lantana *Lantana) error {
	return fmt.Errorf("not implements add lantana lantana.LantanaReps.AddLantana")
}

func (l LantanaReps) SearchLantana(ctx context.Context, query *LantanaSearchQuery) ([]*Lantana, error) {
	lantanas := map[string]*Lantana{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan []*Lantana, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			matchLantana, err := rep.SearchLantana(ctx, query)
			if err != nil {
				errch <- err
				return
			}
			ch <- matchLantana
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at search lantana: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return nil, err
	}

loop:
	for {
		select {
		case matchLantanas := <-ch:
			if matchLantanas == nil {
				continue loop
			}
			for _, matchLantana := range matchLantanas {
				lantanas[matchLantana.LantanaID] = matchLantana
			}
		default:
			break loop
		}
	}

	allLantanas := []*Lantana{}
	for _, lantana := range lantanas {
		if lantana == nil {
			continue
		}
		allLantanas = append(allLantanas, lantana)
	}
	return allLantanas, nil
}

func (l LantanaReps) GetAllKyous(ctx context.Context) ([]*kyou.Kyou, error) {
	kyous := map[string]*kyou.Kyou{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan []*kyou.Kyou, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			matchKyous, err := rep.GetAllKyous(ctx)
			if err != nil {
				errch <- err
				return
			}
			ch <- matchKyous
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at get all kyous: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return nil, err
	}

loop:
	for {
		select {
		case k := <-ch:
			if k == nil {
				continue loop
			}
			for _, kyou := range k {
				kyous[kyou.ID] = kyou
			}
		default:
			break loop
		}
	}

	allKyous := []*kyou.Kyou{}
	for _, kyou := range kyous {
		if kyou == nil {
			continue
		}
		allKyous = append(allKyous, kyou)
	}
	return allKyous, nil
}

func (l LantanaReps) GetContentHTML(ctx context.Context, id string) (string, error) {
	contentHTMLs := []string{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan string, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			contentHTML, err := rep.GetContentHTML(ctx, id)
			if err != nil {
				// errch <- err
				return
			}
			ch <- contentHTML
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at get all kyous: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return "", err
	}

loop:
	for {
		select {
		case h := <-ch:
			if h == "" {
				continue loop
			}
			contentHTMLs = append(contentHTMLs, h)
		default:
			break loop
		}
	}
	for _, contentHTML := range contentHTMLs {
		return contentHTML, nil
	}
	return "", fmt.Errorf("not found content html %s", id)
}

func (l LantanaReps) GetPath(ctx context.Context, id string) (string, error) {
	return "", fmt.Errorf("not implements get path. lantana.LantanaReps.GetPath")
}

func (l LantanaReps) Delete(id string) error {
	return fmt.Errorf("not implements delete. lantana.LantanaReps.Delete")
}

func (l LantanaReps) Close() error {
	var err error
	for _, rep := range l {
		e := rep.Close()
		if e != nil {
			err = fmt.Errorf(":%w", e)
		}
	}
	return err
}

func (l LantanaReps) Path() string {
	return "LantanaReps"
}

func (l LantanaReps) RepName() string {
	return "LantanaReps"
}

func (l LantanaReps) Search(ctx context.Context, word string) ([]*kyou.Kyou, error) {
	kyous := map[string]*kyou.Kyou{}
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan []*kyou.Kyou, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			matchKyous, err := rep.Search(ctx, word)
			if err != nil {
				errch <- err
				return
			}
			ch <- matchKyous
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at get search: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return nil, err
	}

loop:
	for {
		select {
		case k := <-ch:
			if k == nil {
				continue loop
			}
			for _, kyou := range k {
				kyous[kyou.ID] = kyou
			}
		default:
			break loop
		}
	}

	allKyous := []*kyou.Kyou{}
	for _, kyou := range kyous {
		if kyou == nil {
			continue
		}
		allKyous = append(allKyous, kyou)
	}
	return allKyous, nil
}

func (l LantanaReps) UpdateCache(ctx context.Context) error {
	existErr := false
	var err error
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, len(l))
	errch := make(chan error, len(l))
	defer close(ch)
	defer close(errch)
	for _, rep := range l {
		wg.Add(1)
		rep := rep
		go func(rep LantanaRep) {
			defer wg.Done()
			err := rep.UpdateCache(ctx)
			if err != nil {
				errch <- err
				return
			}
			ch <- nil
		}(rep)
	}
	wg.Wait()
errloop:
	for {
		select {
		case e := <-errch:
			err = fmt.Errorf("error at update cache: %w", e)
			existErr = true
		default:
			break errloop
		}
	}
	if existErr {
		return err
	}

loop:
	for {
		select {
		case t := <-ch:
			if t == nil {
				continue loop
			}
		default:
			break loop
		}
	}
	return nil
}
