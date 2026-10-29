
package instructions

var A []int
var B []int

func Push (from *[]int, to *[]int) {
	if len(*from) > 0 {
		*to = append([]int{(*from)[0]}, *to...)
		*from = (*from)[1:]
	}
}

func swap (list []int) {
	if len(list) > 1 {
		list[0], list[1] = list[1], list[0]
	}
}
func ss() {
	swap (A)
	swap (B)
}
func rotate (list *[]int) {
	if len(*list) > 1 {
		*list = append((*list)[1:], (*list)[0])
	}
}
func rr() {
	rotate(&A)
	rotate(&B)
}

func reverseRotate (list *[]int) {
	if len(*list) > 1 {
		*list = append([]int{(*list)[len(*list)-1]}, (*list)[:len(*list)-1]...)
	}
}
func rrr() {
	reverseRotate (&A)
	reverseRotate (&B)
}

