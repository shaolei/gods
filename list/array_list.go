package list

import "fmt"

const initCap = 1

type ArrayList struct {
	data []any
	size int
}

// NewArrayList 默认初始化，容量为1
func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]any, initCap),
		size: 0,
	}
}

// NewArrayListWithCap 按指定容量初始化
func NewArrayListWithCap(cap int) *ArrayList {
	return &ArrayList{
		data: make([]any, cap),
		size: 0,
	}
}

// AddLast 在链表末尾增加元素
func (al *ArrayList) AddLast(e any) {
	capVal := cap(al.data)
	if al.size == capVal {
		al.resize(2 * capVal)
	}
	al.data[al.size] = e
	al.size++
}

// Add 在指定索引增加元素
func (al *ArrayList) Add(index int, e any) error {
	if err := al.checkPositionIndex(index); err != nil {
		return err
	}
	capVal := cap(al.data)
	if al.size == capVal {
		al.resize(2 * capVal)
	}
	copy(al.data[index+1:], al.data[index:])
	al.data[index] = e
	al.size++
	return nil
}

// AddFirst 在链表头增加元素
func (al *ArrayList) AddFirst(e any) error {
	return al.Add(0, e)
}

// RemoveLast 在链表末尾增加元素
func (al *ArrayList) RemoveLast() (any, error) {
	if al.IsEmpty() {
		return nil, fmt.Errorf("no such element")
	}
	capVal := cap(al.data)
	if al.size == capVal/4 {
		al.resize(capVal / 2)
	}
	e := al.data[al.size-1]
	al.data[al.size-1] = nil
	al.size--

	return e, nil
}

// Remove 在指定索引增加元素
func (al *ArrayList) Remove(index int) (any, error) {
	if al.IsEmpty() {
		return nil, fmt.Errorf("no such element")
	}

	if al.size == cap(al.data)/4 {
		al.resize(cap(al.data) / 2)
	}
	old := al.data[index]
	copy(al.data[index:], al.data[index+1:])
	al.size--
	return old, nil
}

// RemoveFirst 在链表头增加元素
func (al *ArrayList) RemoveFirst() (any, error) {
	return al.Remove(0)
}

// Get 获取元素
func (al *ArrayList) Get(index int) (any, error) {
	if err := al.checkElementIndex(index); err != nil {
		return nil, err
	}
	return al.data[index], nil
}

// Set 在指定索引修改元素
func (al *ArrayList) Set(index int, e any) (any, error) {
	if err := al.checkElementIndex(index); err != nil {
		return nil, err
	}
	old := al.data[index]
	al.data[index] = e
	return old, nil
}

//***********工具函数***********//

// Size 返回链表当前元素个数
func (al *ArrayList) Size() int {
	return al.size
}

// IsEmpty 判断链表是否为空
func (al *ArrayList) IsEmpty() bool {
	return al.size == 0
}

//***********私有函数***********//

// 判断索引位置是否存在元素
func (al *ArrayList) isElementIndex(index int) bool {
	return index >= 0 && index < al.size
}

// 判断索引位置是否可以添加元素
func (al *ArrayList) isPositionIndex(index int) bool {
	return index >= 0 && index <= al.size
}

// 检查index索引是否可以存在元素
func (al *ArrayList) checkElementIndex(index int) error {
	if !al.isElementIndex(index) {
		return fmt.Errorf("index out of bounds. index: %d, size: %d", index, al.size)
	}
	return nil
}

// 检查index索引是否可以添加元素
func (al *ArrayList) checkPositionIndex(index int) error {
	if !al.isPositionIndex(index) {
		return fmt.Errorf("index out of bounds. index: %d, size: %d", index, al.size)
	}
	return nil
}

// 更改链表容量
func (al *ArrayList) resize(cap int) {
	if al.size > cap {
		return
	}
	temp := make([]any, cap)
	//for i, v := range al.data {
	//	temp[i] = v
	//}
	copy(temp, al.data)
	al.data = temp
}
