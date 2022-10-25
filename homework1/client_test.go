package homework1

import "testing"

func TestSortClients(t *testing.T) {
	TestCases := []struct{
		id int
		name string
		age int
	}{
		{1, "zohid", 18},
		{2, "alimadad", 19},
		{3, "kamila", 13},
		{4, "olim", 9},
		{5, "abror", 23},
		{6, "rustam", 32},
		{7, "qobil", 11},
		{8, "elnora", 20},
		{9, "vanella", 45},
		{10, "redma", 34},
		{11, "bunyod", 14},
		{12, "bilol", 54},
		{13, "abdurahmon", 21},
		{14, "hilola", 18},
		{15, "diyora", 17},
		{16, "xumoyin", 12},
		{17, "maqsad", 30},
		{18, "lola", 23},
		{19, "karim", 10},
		{20, "zebo", 28},
	}
	for _, testcase := range TestCases {
		if res := SortClients(testcase.id, testcase.name, testcase.age); res != nil {
			t.Errorf("The Test Case -> %v FAILED, Expected %v, Got Error -> %v\n", testcase, nil, res)
			t.Fatal()
		}
	}
}