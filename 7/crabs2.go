package main

import (
	"fmt"
	"math"
)

// crabs burn fuel at a rate of 1 more than the last move. 1, 2, 3, etc
// much more expensive to move further away.
//

func calcFuel(dist int) int {
	sum := 0
	for i := 0; i < dist; i++ {
		sum += i + 1
	}
	return sum
}

func main() {
	crabs := []int{1101, 1, 29, 67, 1102, 0, 1, 65, 1008, 65, 35, 66, 1005, 66, 28, 1, 67, 65, 20, 4, 0, 1001, 65, 1, 65, 1106, 0, 8, 99, 35, 67, 101, 99, 105, 32, 110, 39, 101, 115, 116, 32, 112, 97, 115, 32, 117, 110, 101, 32, 105, 110, 116, 99, 111, 100, 101, 32, 112, 114, 111, 103, 114, 97, 109, 10, 52, 1088, 462, 1398, 576, 241, 636, 512, 28, 390, 168, 262, 6, 489, 1152, 466, 539, 133, 159, 1481, 128, 198, 858, 57, 12, 1155, 400, 137, 557, 1370, 440, 885, 1433, 360, 387, 5, 173, 397, 465, 426, 365, 470, 456, 45, 1052, 1116, 26, 17, 585, 647, 357, 786, 313, 1124, 346, 694, 941, 124, 825, 243, 852, 76, 618, 436, 596, 14, 958, 969, 895, 1745, 246, 822, 239, 952, 928, 206, 406, 190, 459, 841, 25, 1087, 299, 962, 15, 1539, 1003, 456, 51, 546, 858, 137, 1214, 110, 936, 975, 1164, 51, 82, 947, 1354, 312, 132, 261, 181, 287, 107, 1411, 332, 930, 60, 1458, 22, 248, 175, 3, 946, 1097, 35, 231, 648, 109, 313, 1061, 163, 1382, 80, 912, 89, 718, 1068, 419, 703, 155, 321, 909, 9, 212, 478, 315, 118, 206, 38, 125, 130, 1391, 229, 8, 44, 571, 432, 24, 283, 0, 941, 422, 251, 686, 578, 154, 123, 489, 86, 1217, 129, 227, 638, 47, 187, 946, 2, 536, 227, 640, 1170, 1444, 286, 1280, 83, 1253, 1735, 70, 52, 104, 658, 367, 302, 462, 394, 13, 798, 514, 104, 260, 479, 526, 632, 1161, 1118, 320, 196, 262, 571, 1319, 594, 131, 797, 37, 566, 1054, 271, 159, 1021, 244, 204, 447, 624, 825, 723, 364, 234, 105, 362, 305, 391, 681, 692, 89, 380, 104, 1217, 814, 1467, 898, 207, 1345, 94, 10, 1380, 50, 1192, 178, 1539, 1712, 145, 390, 9, 878, 144, 1241, 395, 10, 41, 80, 1719, 1077, 113, 46, 1699, 130, 91, 723, 359, 1617, 1065, 530, 1058, 903, 163, 412, 45, 858, 10, 1704, 141, 451, 1314, 879, 13, 857, 905, 87, 830, 1228, 25, 1594, 153, 4, 585, 46, 862, 265, 833, 301, 473, 458, 85, 254, 22, 266, 543, 32, 939, 1113, 228, 544, 205, 1617, 1109, 445, 86, 5, 278, 16, 784, 303, 1022, 1014, 162, 714, 447, 656, 834, 138, 448, 30, 85, 371, 951, 1256, 842, 5, 460, 919, 1019, 785, 1275, 616, 1593, 168, 1727, 311, 950, 1299, 1131, 796, 522, 443, 703, 836, 47, 300, 449, 11, 360, 682, 487, 108, 1396, 623, 1108, 239, 379, 0, 822, 109, 60, 98, 667, 242, 1398, 650, 25, 376, 168, 46, 259, 138, 254, 1631, 953, 776, 166, 0, 628, 75, 413, 1401, 69, 462, 883, 877, 96, 314, 825, 346, 932, 352, 1086, 143, 507, 134, 557, 31, 1663, 565, 275, 207, 330, 702, 53, 1085, 259, 14, 26, 851, 1571, 1829, 1513, 356, 70, 1393, 426, 345, 412, 129, 908, 959, 896, 1578, 617, 428, 222, 1256, 3, 863, 237, 5, 357, 92, 292, 514, 4, 919, 5, 848, 1605, 149, 959, 376, 1709, 410, 460, 646, 389, 13, 1388, 294, 1151, 652, 10, 113, 769, 1519, 57, 685, 1132, 1, 417, 1369, 1396, 248, 496, 145, 64, 798, 719, 716, 845, 168, 1, 147, 347, 239, 512, 19, 478, 336, 20, 327, 487, 141, 37, 96, 331, 826, 1347, 479, 182, 601, 233, 564, 196, 0, 811, 19, 318, 86, 1442, 468, 396, 298, 46, 661, 339, 914, 54, 560, 91, 284, 829, 1710, 478, 318, 780, 738, 807, 1017, 166, 48, 358, 193, 466, 831, 138, 226, 5, 24, 251, 119, 644, 545, 588, 170, 890, 248, 596, 310, 612, 479, 366, 1374, 465, 32, 467, 79, 603, 220, 1138, 168, 968, 420, 129, 90, 214, 652, 408, 169, 0, 173, 19, 312, 65, 38, 115, 325, 158, 1458, 744, 1529, 361, 360, 77, 75, 130, 111, 175, 34, 676, 169, 384, 473, 296, 701, 84, 11, 1862, 223, 193, 118, 678, 403, 1097, 2, 1318, 190, 590, 96, 47, 69, 212, 520, 786, 1569, 703, 1776, 140, 12, 741, 906, 29, 115, 30, 196, 821, 23, 51, 540, 225, 891, 133, 907, 567, 143, 44, 371, 1038, 237, 0, 222, 1327, 760, 854, 1, 29, 14, 65, 98, 25, 233, 423, 63, 382, 648, 257, 160, 71, 1287, 315, 627, 40, 159, 202, 112, 657, 87, 94, 93, 362, 23, 501, 870, 1114, 946, 1007, 453, 159, 493, 590, 665, 28, 435, 7, 1238, 1846, 758, 174, 258, 972, 557, 1431, 482, 429, 57, 389, 651, 1089, 1490, 821, 844, 458, 712, 259, 433, 418, 344, 466, 60, 123, 1604, 897, 1346, 198, 143, 259, 49, 770, 1703, 900, 1364, 450, 498, 30, 543, 322, 3, 533, 508, 444, 148, 927, 72, 321, 733, 689, 24, 44, 685, 1021, 324, 182, 1737, 975, 387, 143, 176, 478, 602, 752, 203, 130, 169, 165, 41, 119, 35, 175, 763, 1147, 5, 137, 10, 357, 54, 1209, 182, 298, 156, 1488, 176, 86, 548, 2, 37, 36, 76, 100, 1369, 1174, 322, 32, 573, 107, 375, 1210, 51, 597, 902, 878, 919, 379, 125, 26, 1240, 7, 7, 131, 913, 994, 1097, 576, 112, 694, 805, 551, 512, 663, 361, 747, 161, 691, 63, 119, 47, 89, 6, 258, 57, 537, 654, 757, 1202, 922, 475, 347, 193, 79, 1177, 443, 33, 1257, 1070, 118, 810, 117, 37, 226, 230, 552, 618, 341, 530, 681, 1015, 358, 846, 276, 1149, 210, 525, 1144, 272, 30, 551, 55, 512, 229, 90, 1144, 389, 500, 372, 92, 58, 598, 1362, 475, 70, 748, 1217, 442, 28, 334, 369, 768, 169, 405, 1058, 759, 1087, 268, 714, 81, 594, 1423, 1004, 694, 61, 1032, 895, 1321, 95, 1512, 646, 818, 845, 1275, 294, 883, 1684, 1062, 2, 851, 304, 306, 128, 1523, 1594, 190, 73, 809, 175, 321, 407, 424, 109, 48, 234, 437, 968, 284, 1069, 181, 340, 149, 9, 163, 863, 17, 584, 421, 79, 164, 913, 81}
	fmt.Println("crabs", crabs)
	sum := 0
	for _, v := range crabs {
		sum += v
	}
	fmt.Println("sum", sum)
	avg := sum / len(crabs)
	// avg := 5
	// avg == 4
	// average distance from mean?
	// var sum2 float64
	// for _, v := range crabs {
	// 	fmt.Println("v - avg", v-avg)
	// 	sum2 += math.Abs(float64(v - avg))
	// }
	// idea: take average
	// calculate fuel +/- 1000 from average
	// pick smallest value
	// this is kind of cheating, but it works
	min := -1
	maxDiff := 100
	for i := avg - maxDiff; i < avg+maxDiff; i++ {
		fuel := 0
		for _, v := range crabs {
			// to get the
			fuel += calcFuel(int(math.Abs(float64(v - i))))
		}
		fmt.Println("i, fuel", i, fuel)
		if fuel < min || min == -1 {
			min = fuel
		}
	}
	fmt.Println("fuel", min)
	// fmt.Println("sum2/len(crabs)", sum2/float64(len(crabs)))
}
