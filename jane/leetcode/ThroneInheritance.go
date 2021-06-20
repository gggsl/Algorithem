package main

type ThroneInheritance struct {
	kingName  string
	childMap  map[string][]string
	deathList map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		kingName:  kingName,
		childMap:  make(map[string][]string),
		deathList: make(map[string]bool),
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	childList, exist := this.childMap[parentName]
	if exist {
		childList = append(childList, childName)
	} else {
		childList = []string{childName}
	}
	this.childMap[parentName] = childList
	//this.childMap[parentName] = append(this.childMap[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.deathList[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	ans := []string{}
	var recursion func(string)
	recursion = func(name string) {
		if !this.deathList[name] {
			ans = append(ans, name)
		}
		for _, childName := range this.childMap[name] {
			recursion(childName)
		}
	}
	recursion(this.kingName)
	return ans
}
