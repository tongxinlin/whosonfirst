package main

import()

	var sL1 *PairTreeNode = NewPairTreeNode("strong-cues")
	var sL2 *PairTreeNode = NewPairTreeNode("sl2")
	var sL3 *PairTreeNode = NewPairTreeNode("sl3")

	var sL1L1 *PairTreeNode = NewPairTreeNode("sl1l1")

	var sL1L1L1 *PairTreeNode = NewPairTreeNode("the name")

	var sL1L1L2 *PairTreeNode = NewPairTreeNode("their names")


	var sL1L2 *PairTreeNode = NewPairTreeNode("sl1l2")

	var sL1L2L1 *PairTreeNode = NewPairTreeNode("whos on first whats on second i dont know on third")

	var sL1L2L2 *PairTreeNode = NewPairTreeNode("whats on second whos on first i dont know on third")


	var sL2L1 *PairTreeNode = NewPairTreeNode("sl2l1")

	var sL2L1L1 *PairTreeNode = NewPairTreeNode("suppose")

	var sL2L1L2 *PairTreeNode = NewPairTreeNode("lets say")

	var sL2L1L3 *PairTreeNode = NewPairTreeNode("assume")


	var sL2L2 *PairTreeNode = NewPairTreeNode("sl2l2")

	var sL2L2L1 *PairTreeNode = NewPairTreeNode("okay")

	var sL2L2L2 *PairTreeNode = NewPairTreeNode("why not")

	var sL2L2L3 *PairTreeNode = NewPairTreeNode("sure")

	var sL2L2L4 *PairTreeNode = NewPairTreeNode("it could happen")


	var sL3L1 *PairTreeNode = NewPairTreeNode("sl3l1")

	var sL3L1L1 *PairTreeNode = NewPairTreeNode("i dont know")

	var sL3L2 *PairTreeNode = NewPairTreeNode("sl3l2")

	var sL3L2L1 *PairTreeNode = NewPairTreeNode("third base")

	var sL3L2L2 *PairTreeNode = NewPairTreeNode("hes on third")


//////////////////////////////////////////////////
	var wL1 *PairTreeNode = NewPairTreeNode("weak-cues")
	var wL2 *PairTreeNode = NewPairTreeNode("wl2")
	var wL3 *PairTreeNode = NewPairTreeNode("wl3")


	var wL1L1 *PairTreeNode = NewPairTreeNode("wl1l1")

	var wL1L1L1 *PairTreeNode = NewPairTreeNode("who")

	var wL1L1L2 *PairTreeNode = NewPairTreeNode("whos")

	var wL1L1L3 *PairTreeNode = NewPairTreeNode("who is")


	var wL1L2 *PairTreeNode = NewPairTreeNode("first-base")

	var wL1L2L1 *PairTreeNode = NewPairTreeNode("thats right")

	var wL1L2L2 *PairTreeNode = NewPairTreeNode("exactly")

	var wL1L2L3 *PairTreeNode = NewPairTreeNode("you got it")

	var wL1L2L4 *PairTreeNode = NewPairTreeNode("right on")

	var wL1L2L5 *PairTreeNode = NewPairTreeNode("now youve got it")

	var wL1L3 *PairTreeNode = NewPairTreeNode("second-base third-base")

	var wL1L3L1 *PairTreeNode = NewPairTreeNode("no whos on first")

	var wL1L3L2 *PairTreeNode = NewPairTreeNode("whos on first")

	var wL1L3L3 *PairTreeNode = NewPairTreeNode("first base")


	var wL2L1 *PairTreeNode = NewPairTreeNode("wl2l1")

	var wL2L1L1 *PairTreeNode = NewPairTreeNode("what")

	var wL2L1L2 *PairTreeNode = NewPairTreeNode("whats")

	var wL2L1L3 *PairTreeNode = NewPairTreeNode("what is")

	var wL2L2 *PairTreeNode = NewPairTreeNode("first-base third-base")

	var wL2L2L1 *PairTreeNode = NewPairTreeNode("hes on second")

	var wL2L2L2 *PairTreeNode = NewPairTreeNode("i told you whats on second")

	var wL2L3 *PairTreeNode = NewPairTreeNode("second-base")

	var wL2L3L1 *PairTreeNode = NewPairTreeNode("right")

	var wL2L3L2 *PairTreeNode = NewPairTreeNode("sure")

	var wL2L3L3 *PairTreeNode = NewPairTreeNode("you got it right")


	var wL3L1 *PairTreeNode= NewPairTreeNode("wl3l1")

	var wL3L1L1 *PairTreeNode = NewPairTreeNode("whats the name")

	var wL3L2 *PairTreeNode = NewPairTreeNode("first-base third-base")

	var wL3L2L1 *PairTreeNode = NewPairTreeNode("no whats the name of the guy on second")

	var wL3L2L2 *PairTreeNode = NewPairTreeNode("whats the name of the second baseman")

	var wL3L3 *PairTreeNode = NewPairTreeNode("second-base")

	var wL3L3L1 *PairTreeNode = NewPairTreeNode("now youre talking")

	var wL3L3L2 *PairTreeNode = NewPairTreeNode("you got it")


//////////////////////////////////////////////////////////

	var cL1 *PairTreeNode = NewPairTreeNode("context-words")
	var cL2 *PairTreeNode = NewPairTreeNode("cl2")
	var cL3 *PairTreeNode = NewPairTreeNode("cl3")


	var cL1L1 *PairTreeNode = NewPairTreeNode("cl1l1")

	var cL1L1L1 *PairTreeNode = NewPairTreeNode("first")

	var cL1L2 *PairTreeNode = NewPairTreeNode("cl1l2")

	var cL1L2L1 *PairTreeNode = NewPairTreeNode("first-base")


	var cL2L1 *PairTreeNode = NewPairTreeNode("cl2l1")

	var cL2L1L1 *PairTreeNode = NewPairTreeNode("second")

	var cL2L2 *PairTreeNode = NewPairTreeNode("cl2l2")

	var cL2L2L1 *PairTreeNode = NewPairTreeNode("second-base")


	var cL3L1 *PairTreeNode = NewPairTreeNode("cl3l1")

	var cL3L1L1 *PairTreeNode = NewPairTreeNode("third")

	var cL3L2 *PairTreeNode = NewPairTreeNode("cl3l2")

	var cL3L2L1 *PairTreeNode = NewPairTreeNode("third-base")


/////////////////////////////////////////////////////////

	var hL1 *PairTreeNode = NewPairTreeNode("its like im telling you")
	var hL2 *PairTreeNode = NewPairTreeNode("now calm down")
	var hL3 *PairTreeNode = NewPairTreeNode("take it easy")
	var hL4 *PairTreeNode = NewPairTreeNode("its elementary lou")
	var hL5 *PairTreeNode = NewPairTreeNode("im trying to tell you")
	var hL6 *PairTreeNode = NewPairTreeNode("but you asked")

func SetCues(){

	sL1.SetLeft(sL2)
	sL2.SetLeft(sL3)


	sL1.SetRight(sL1L1)

	sL1L1.SetRight(sL1L1L1)

	sL1L1L1.SetLeft(sL1L1L2)


	sL1L1.SetLeft(sL1L2)

	sL1L2.SetRight(sL1L2L1)

	sL1L2L1.SetLeft(sL1L2L2)


	sL2.SetRight(sL2L1)

	sL2L1.SetRight(sL2L1L1)

	sL2L1L1.SetLeft(sL2L1L2)

	sL2L1L2.SetLeft(sL2L1L3)


	sL2L1.SetLeft(sL2L2)

	sL2L2.SetRight(sL2L2L1)

	sL2L2L1.SetLeft(sL2L2L2)

	sL2L2L2.SetLeft(sL2L2L3)

	sL2L2L3.SetLeft(sL2L2L4)


	sL3.SetRight(sL3L1)

	sL3L1.SetRight(sL3L1L1)

	sL3L1.SetLeft(sL3L2)

	sL3L2.SetRight(sL3L2L1)

	sL3L2L1.SetLeft(sL3L2L2)

//////////////////////////////////////////////////

	wL1.SetLeft(wL2)
	wL2.SetLeft(wL3)


	wL1.SetRight(wL1L1)

	wL1L1.SetRight(wL1L1L1)

	wL1L1L1.SetLeft(wL1L1L2)

	wL1L1L2.SetLeft(wL1L1L3)

	wL1L1.SetLeft(wL1L2)

	wL1L2.SetRight(wL1L2L1)

	wL1L2L1.SetLeft(wL1L2L2)

	wL1L2L2.SetLeft(wL1L2L3)

	wL1L2L3.SetLeft(wL1L2L4)

	wL1L2L4.SetLeft(wL1L2L5)

	wL1L2.SetLeft(wL1L3)

	wL1L3.SetRight(wL1L3L1)

	wL1L3L1.SetLeft(wL1L3L2)

	wL1L3L2.SetLeft(wL1L3L3)

	wL2.SetRight(wL2L1)

	wL2L1.SetRight(wL2L1L1)

	wL2L1L1.SetLeft(wL2L1L2)

	wL2L1L2.SetLeft(wL2L1L3)

	wL2L1.SetLeft(wL2L2)

	wL2L2.SetRight(wL2L2L1)
	
	wL2L2L1.SetLeft(wL2L2L2)

	wL2L2.SetLeft(wL2L3)

	wL2L3.SetRight(wL2L3L1)

	wL2L3L1.SetLeft(wL2L3L2)

	wL2L3L2.SetLeft(wL2L3L3)


	wL3.SetRight(wL3L1)

	wL3L1.SetRight(wL3L1L1)

	wL3L1.SetLeft(wL3L2)

	wL3L2.SetRight(wL3L2L1)

	wL3L2L1.SetLeft(wL3L2L2)

	wL3L2.SetLeft(wL3L3)

	wL3L3.SetRight(wL3L3L1)

	wL3L3L1.SetLeft(wL3L3L2)

//////////////////////////////////////////////////////////

	cL1.SetLeft(cL2)
	cL2.SetLeft(cL3)


	cL1.SetRight(cL1L1)

	cL1L1.SetRight(cL1L1L1)

	cL1L1.SetLeft(cL1L2)

	cL1L2.SetRight(cL1L2L1)


	cL2.SetRight(cL2L1)

	cL2L1.SetRight(cL2L1L1)

	cL2L1.SetLeft(cL2L2)

	cL2L2.SetRight(cL2L2L1)


	cL3.SetRight(cL3L1)

	cL3L1.SetRight(cL3L1L1)

	cL3L1.SetLeft(cL3L2)

	cL3L2.SetRight(cL3L2L1)

/////////////////////////////////////////////////////////

	hL1.SetLeft(hL2)
	hL2.SetLeft(hL3)
	hL3.SetLeft(hL4)
	hL4.SetLeft(hL5)
	hL5.SetLeft(hL6)
}