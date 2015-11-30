package main

import(
	"math/rand"
	"time"
	"strings"
)

func (this *PairTreeNode) MyListRef(pos int) string {
	if pos <= 0 || pos > this.GetLength() {
		return ""
	}
	var i int = 1
	var t *PairTreeNode = this
	for i < pos {
		//t = t.GetLeft()
		t = t.Cdr()
		i++
	}
	return t.GetValue()
}

func Random(n int) int{
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return r.Intn(n) + 1
}

func SelectAnyFromList(t *PairTreeNode) string {
	return t.MyListRef(Random(t.GetLength()))
}

func IsSubsequence(e string, sentence string) bool {
	return strings.Contains(sentence,e)
}

func AnyGoodFragments(listOfCues *PairTreeNode, sentence string) bool {
	if listOfCues == nil {
		return false
	} else if IsSubsequence(listOfCues.GetValue(), sentence) {
		return true
	} else {
		return AnyGoodFragments(listOfCues.Cdr(), sentence)
	}
}

func TryStrongCuesHelper(sentence string, listOfCues *PairTreeNode) string {
	if listOfCues == nil {
		return "nil"
	} else if AnyGoodFragments(listOfCues.Car().Car(), sentence) {
		return SelectAnyFromList(listOfCues.Car().Cdr().Car())
	} else {
		return TryStrongCuesHelper(sentence,listOfCues.Cdr())
	}
}

func Hedge(listOfCues *PairTreeNode) string {
	return SelectAnyFromList(listOfCues)
}

func GetContextHelper(sentence string, oldContext string, listOfCues *PairTreeNode) string {
	if listOfCues == nil {
		return oldContext
	} else if AnyGoodFragments(listOfCues.Car().Car(), sentence) {
		return listOfCues.Car().Cdr().Car().GetValue()
	} else {
		return GetContextHelper(sentence, oldContext, listOfCues.Cdr())
	}
}

func IsExist(item string, list string) bool {
	if item == "" {
		return false
	}
	return strings.Contains(list,item)
}

func TryWeakCuesHelper(sentence string, context string, listOfCues *PairTreeNode) string {
	if listOfCues == nil {
		return "nil"
	} else if AnyGoodFragments(listOfCues.Car().Car(),sentence) {
		var t *PairTreeNode = listOfCues.Car().Cdr()
		if IsExist(context, t.GetValue()) {
			return SelectAnyFromList(t.Car())
		} else {
			for t.HasLeft() {
				t = t.Cdr()
				if IsExist(context, t.GetValue()) {
					return SelectAnyFromList(t.Car())
				}
			}
			return "nil"
		}
	} else {
		return TryWeakCuesHelper(sentence, context, listOfCues.Cdr())
	}
}

var context string = ""

func WhosOnFirst(oldContext string, readin string, strongCues *PairTreeNode, weakCues *PairTreeNode, contextWords *PairTreeNode, hedges *PairTreeNode) string{
	var costellosLine string = readin
	var newContext string = GetContextHelper(costellosLine, oldContext, contextWords)
	var strongReply string = TryStrongCuesHelper(costellosLine, strongCues)
	var weakReply string = TryWeakCuesHelper(costellosLine, newContext, weakCues)
	if readin == "i quit<br/>" {
		context = ""
		return "have a nice day"
	} else if strongReply != "nil" {
		context = GetContextHelper(strongReply, newContext, contextWords)
		return strongReply
	} else if weakReply != "nil" {
		context = GetContextHelper(weakReply, newContext, contextWords)
		return weakReply
	} else {
		context = newContext
		return Hedge(hedges)
	}
}