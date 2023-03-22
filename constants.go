package main

var defaultTaskUrl string = "http://ceko-pmr.org/uploads/ege/подготовка к егэ"
var defaultTaskUrlMock string = "/банк заданий егэ/ооо с русским языком обучения"

var subjects = map[string]string{
	"math": "математика",
}

var TasksOfBlock1Math = [][]string{
	{"1","2", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_012.jpg"},
	{"3","15", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_013.jpg"},
	{"16","29", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_014.jpg"},
	{"30","42", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_015.jpg"},
	{"43","55", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_016.jpg"},
	{"56","68", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_017.jpg"},
	{"69","81", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_018.jpg"},
	{"82","94", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_019.jpg"},
	{"95","107", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_020.jpg"},
	{"108","121", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_021.jpg"},
	{"122","134", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_022.jpg"},
	{"135","145", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 01/математика р_страница_023.jpg"},
}

var TasksOfBlock2Math = [][]string{
	{"1","4", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_023.jpg"},
	{"5","17", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_024.jpg"},
	{"18","30", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_025.jpg"},
	{"31","44", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_026.jpg"},
	{"45","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_027.jpg"},
	{"58","68", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_028.jpg"},
	{"69","82", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_029.jpg"},
	{"83","96", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_030.jpg"},
	{"97","110", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_031.jpg"},
	{"111","122", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_032.jpg"},
	{"123","135", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_033.jpg"},
	{"136","145", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 02/математика р_страница_034.jpg"},
}

var TasksOfBlock3Math = [][]string{
	{"1","1", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_034.jpg"},
	{"2","6", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_035.jpg"},
	{"7","10", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_036.jpg"},
	{"11","15", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_037.jpg"},
	{"16","20", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_038.jpg"},
	{"21","26", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_039.jpg"},
	{"27","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_040.jpg"},
	{"32","36", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_041.jpg"},
	{"37","41", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_042.jpg"},
	{"42","46", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_043.jpg"},
	{"47","51", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_044.jpg"},
	{"52","56", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_045.jpg"},
	{"57","61", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_046.jpg"},
	{"62","66", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_047.jpg"},
	{"67","71", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_048.jpg"},
	{"72","77", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_049.jpg"},
	{"78","83", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_050.jpg"},
	{"84","88", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_051.jpg"},
	{"89","93", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_052.jpg"},
	{"94","98", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_053.jpg"},
	{"99","104", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_054.jpg"},
	{"105","110", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_55.jpg"},
	{"111","115", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_056.jpg"},
	{"116","120", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_057jpg"},
	{"121","125", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_058.jpg"},
	{"126","130", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_059.jpg"},
	{"131","137", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_060.jpg"},
	{"138","140", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 03/математика р_страница_061.jpg"},
}

var TasksOfBlock4Math = [][]string{
	{"1","6", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_061.jpg"},
	{"7","20", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_062.jpg"},
	{"21","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_063.jpg"},
	{"32","44", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_064.jpg"},
	{"45","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_065.jpg"},
	{"58","70", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_066.jpg"},
	{"71","83", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_067.jpg"},
	{"84","96", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_068.jpg"},
	{"97","111", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_069.jpg"},
	{"112","123", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_070.jpg"},
	{"124","136", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_071.jpg"},
	{"137","150", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_072.jpg"},
	{"151","162", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_073.jpg"},
	{"163","170", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 04/математика р_страница_074.jpg"},
}

var TasksOfBlock5Math = [][]string{
	{"1","2", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_074.jpg"},
	{"3","10", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_075.jpg"},
	{"11","19", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_076.jpg"},
	{"20","27", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_077.jpg"},
	{"28","35", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_078.jpg"},
	{"36","43", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_079.jpg"},
	{"44","50", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_080.jpg"},
	{"51","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_081.jpg"},
	{"58","64", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_082.jpg"},
	{"65","72", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_083.jpg"},
	{"73","80", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_084.jpg"},
	{"81","89", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_085.jpg"},
	{"90","97", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_086.jpg"},
	{"98","105", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_087.jpg"},
	{"106","115", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_088.jpg"},
	{"116","124", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_089.jpg"},
	{"125","132", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_090.jpg"},
	{"133","139", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_091.jpg"},
	{"140","147", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_092.jpg"},
	{"148","153", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_093.jpg"},
	{"154","159", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_094.jpg"},
	{"160","167", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_095.jpg"},
	{"168","175", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_096.jpg"},
	{"176","177", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 05/математика р_страница_097.jpg"},
}

var TasksOfBlock6Math = [][]string{
	{"1","13", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_097.jpg"},
	{"14","30", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_098.jpg"},
	{"31","44", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_099.jpg"},
	{"45","56", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_100.jpg"},
	{"57","71", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_101.jpg"},
	{"72","87", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_102.jpg"},
	{"88","100", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_103.jpg"},
	{"101","114", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_104.jpg"},
	{"115","128", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_105.jpg"},
	{"129","145", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_106.jpg"},
	{"146","159", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_107.jpg"},
	{"160","160", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 06/математика р_страница_108.jpg"},
}

var TasksOfBlock7Math = [][]string{
	{"1","16", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_108.jpg"},
	{"17","33", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_109.jpg"},
	{"34","54", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_110.jpg"},
	{"55","75", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_111.jpg"},
	{"76","93", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_112.jpg"},
	{"94","115", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_113.jpg"},
	{"116","134", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_114.jpg"},
	{"135","151", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_115.jpg"},
	{"152","164", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_116.jpg"},
	{"165","170", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 07/математика р_страница_117.jpg"},
}

var TasksOfBlock8Math = [][]string{
	{"1","8", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_117.jpg"},
	{"9","17", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_118.jpg"},
	{"18","27", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_119.jpg"},
	{"28","37", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_120.jpg"},
	{"38","46", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_121.jpg"},
	{"47","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_122.jpg"},
	{"58","67", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_123.jpg"},
	{"68","78", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_124.jpg"},
	{"79","87", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_125.jpg"},
	{"88","97", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_126.jpg"},
	{"98","107", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_127.jpg"},
	{"108","118", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_128.jpg"},
	{"119","128", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_129.jpg"},
	{"129","130", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 08/математика р_страница_130.jpg"},
}

var TasksOfBlock9Math = [][]string{
	{"1","3", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_130.jpg"},
	{"4","7", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_131.jpg"},
	{"8","11", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_132.jpg"},
	{"12","14", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_133.jpg"},
	{"15","17", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_134.jpg"},
	{"18","20", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_135.jpg"},
	{"21","24", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_136.jpg"},
	{"25","28", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_137.jpg"},
	{"29","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_138.jpg"},
	{"32","35", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_139.jpg"},
	{"36","38", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_140.jpg"},
	{"39","41", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_141.jpg"},
	{"42","43", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_142.jpg"},
	{"44","46", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_143.jpg"},
	{"47","49", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_144.jpg"},
	{"50","53", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_145.jpg"},
	{"54","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_146.jpg"},
	{"58","61", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_147.jpg"},
	{"62","64", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_148.jpg"},
	{"65","67", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_149.jpg"},
	{"68","70", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_150.jpg"},
	{"71","73", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_151.jpg"},
	{"74","76", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_152.jpg"},
	{"77","80", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_153.jpg"},
	{"81","83", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_154.jpg"},
	{"84","87", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_155.jpg"},
	{"88","90", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_156.jpg"},
	{"91","93", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_157.jpg"},
	{"94","96", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_158.jpg"},
	{"97","99", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_159.jpg"},
	{"100","103", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_160.jpg"},
	{"104","107", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_161.jpg"},
	{"108","110", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_162.jpg"},
	{"111","114", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_163.jpg"},
	{"115","117", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_164.jpg"},
	{"118","120", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_165.jpg"},
	{"121","122", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_166.jpg"},
	{"123","125", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_167.jpg"},
	{"126","128", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_168.jpg"},
	{"129","130", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 09/математика р_страница_169.jpg"},
}

var TasksOfBlock10Math = [][]string{
	{"1","2", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_169.jpg"},
	{"3","9", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_170.jpg"},
	{"10","16", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_171.jpg"},
	{"17","24", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_172.jpg"},
	{"25","33", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_173.jpg"},
	{"34","42", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_174.jpg"},
	{"43","54", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_175.jpg"},
	{"55","63", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_176.jpg"},
	{"64","73", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_177.jpg"},
	{"74","84", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_178.jpg"},
	{"85","94", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_179.jpg"},
	{"95","104", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_180.jpg"},
	{"105","114", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_181.jpg"},
	{"115","124", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_182.jpg"},
	{"125","134", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_183.jpg"},
	{"135","146", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_184.jpg"},
	{"147","150", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 10/математика р_страница_185.jpg"},
}

var TasksOfBlock11Math = [][]string{
	{"1","6", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_185.jpg"},
	{"7","17", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_186.jpg"},
	{"18","22", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_187.jpg"},
	{"23","27", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_188.jpg"},
	{"28","32", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_189.jpg"},
	{"33","37", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_190.jpg"},
	{"38","42", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_191.jpg"},
	{"43","47", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_192.jpg"},
	{"48","52", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_193.jpg"},
	{"53","57", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_194.jpg"},
	{"58","62", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_195.jpg"},
	{"63","69", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_196.jpg"},
	{"70","74", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_197.jpg"},
	{"75","78", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_198.jpg"},
	{"79","84", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_199.jpg"},
	{"85","89", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_200.jpg"},
	{"90","92", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_201.jpg"},
	{"93","95", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_202.jpg"},
	{"96","100", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_203.jpg"},
	{"101","103", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_204.jpg"},
	{"104","107", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_205.jpg"},
	{"108","110", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_206.jpg"},
	{"111","114", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_207.jpg"},
	{"115","117", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_208.jpg"},
	{"118","120", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_209.jpg"},
	{"121","124", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_210.jpg"},
	{"125","128", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_211.jpg"},
	{"129","132", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_212.jpg"},
	{"133","135", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_213.jpg"},
	{"136","139", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_214.jpg"},
	{"140","140", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 11/математика р_страница_215.jpg"},
}

var TasksOfBlock12Math = [][]string{
	{"1","3", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_215.jpg"},
	{"4","7", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_216.jpg"},
	{"8","10", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_217.jpg"},
	{"11","14", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_218.jpg"},
	{"15","18", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_219.jpg"},
	{"19","22", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_220.jpg"},
	{"23","25", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_221.jpg"},
	{"26","28", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_222.jpg"},
	{"29","32", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_223.jpg"},
	{"33","36", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_224.jpg"},
	{"37","41", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_225.jpg"},
	{"42","44", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_226.jpg"},
	{"45","48", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_227.jpg"},
	{"49","52", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_228.jpg"},
	{"53","55", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_229.jpg"},
	{"56","60", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_230.jpg"},
	{"61","64", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_231.jpg"},
	{"65","68", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_232.jpg"},
	{"69","72", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_233.jpg"},
	{"73","75", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_234.jpg"},
	{"76","79", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_235.jpg"},
	{"80","83", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_236.jpg"},
	{"84","87", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_237.jpg"},
	{"88","91", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_238.jpg"},
	{"92","95", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_239.jpg"},
	{"96","100", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 12/математика р_страница_240.jpg"},
}

var TasksOfBlock13Math = [][]string{
	{"1","15", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_241.jpg"},
	{"16","30", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_242.jpg"},
	{"31","49", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_243.jpg"},
	{"50","67", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_244.jpg"},
	{"68","85", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_245.jpg"},
	{"86","102", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_246.jpg"},
	{"103","118", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_247.jpg"},
	{"119","132", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_248.jpg"},
	{"133","146", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_249.jpg"},
	{"147","159", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 13/математика р_страница_250.jpg"},
}

var TasksOfBlock14Math = [][]string{
	{"1","8", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_251.jpg"},
	{"9","18", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_252.jpg"},
	{"19","28", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_253.jpg"},
	{"29","38", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_254.jpg"},
	{"39","50", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_255.jpg"},
	{"51","61", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_256.jpg"},
	{"62","72", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_257.jpg"},
	{"73","84", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_258.jpg"},
	{"85","94", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_259.jpg"},
	{"95","104", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_260.jpg"},
	{"105","113", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_261.jpg"},
	{"114","123", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_262.jpg"},
	{"124","132", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_263.jpg"},
	{"133","135", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 14/математика р_страница_264.jpg"},
}

var TasksOfBlock15Math = [][]string{
	{"1","13", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_264.jpg"},
	{"14","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_265.jpg"},
	{"32","45", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_266.jpg"},
	{"46","58", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_267.jpg"},
	{"59","78", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_268.jpg"},
	{"79","95", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_269.jpg"},
	{"96","112", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_270.jpg"},
	{"113","128", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_271.jpg"},
	{"129","140", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 15/математика р_страница_272.jpg"},
}

var TasksOfBlock16Math = [][]string{
	{"1","4", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_272.jpg"},
	{"5","18", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_273.jpg"},
	{"19","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_274.jpg"},
	{"32","43", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_275.jpg"},
	{"44","56", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_276.jpg"},
	{"57","70", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_277.jpg"},
	{"71","86", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_278.jpg"},
	{"87","101", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_279.jpg"},
	{"102","113", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_280.jpg"},
	{"114","127", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_281.jpg"},
	{"128","139", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_282.jpg"},
	{"140","140", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 16/математика р_страница_283.jpg"},
}

var TasksOfBlock17Math = [][]string{
	{"1","11", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_283.jpg"},
	{"12","23", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_284.jpg"},
	{"24","34", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_285.jpg"},
	{"35","46", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_286.jpg"},
	{"47","58", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_287.jpg"},
	{"59","72", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_288.jpg"},
	{"73","84", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_289.jpg"},
	{"85","95", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_290.jpg"},
	{"96","107", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_291.jpg"},
	{"108","118", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_292.jpg"},
	{"119","125", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 17/математика р_страница_293.jpg"},
}

var TasksOfBlock18Math = [][]string{
	{"1","6", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 18/математика р_страница_293.jpg"},
	{"7","27", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 18/математика р_страница_294.jpg"},
	{"28","48", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 18/математика р_страница_295.jpg"},
	{"49","66", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 18/математика р_страница_296.jpg"},
	{"67","80", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 18/математика р_страница_297.jpg"},
}

var TasksOfBlock19Math = [][]string{
	{"1","2", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_297.jpg"},
	{"3","14", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_298.jpg"},
	{"15","27", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_299.jpg"},
	{"28","40", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_300.jpg"},
	{"41","53", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_301.jpg"},
	{"54","65", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_302.jpg"},
	{"66","78", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_303.jpg"},
	{"79","84", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 19/математика р_страница_304.jpg"},
}

var TasksOfBlock20Math = [][]string{
	{"1","5", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_304.jpg"},
	{"6","17", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_305.jpg"},
	{"18","31", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_306.jpg"},
	{"32","43", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_307.jpg"},
	{"44","55", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_308.jpg"},
	{"56","68", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_309.jpg"},
	{"69","77", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 20/математика р_страница_310.jpg"},
}

var TasksOfBlock21Math = [][]string{
	{"1","5", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_310.jpg"},
	{"6","21", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_311.jpg"},
	{"22","37", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_312.jpg"},
	{"38","51", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_313.jpg"},
	{"52","65", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_314.jpg"},
	{"66","79", defaultTaskUrl+"/"+subjects["math"]+defaultTaskUrlMock+"/задание 21/математика р_страница_315.jpg"},
}

var mathBlocks = [][][]string{
	TasksOfBlock1Math,
	TasksOfBlock2Math,
	TasksOfBlock3Math,
	TasksOfBlock4Math,
	TasksOfBlock5Math,
	TasksOfBlock6Math,
	TasksOfBlock7Math,
	TasksOfBlock8Math,
	TasksOfBlock9Math,
	TasksOfBlock10Math,
	TasksOfBlock11Math,
	TasksOfBlock12Math,
	TasksOfBlock13Math,
	TasksOfBlock14Math,
	TasksOfBlock15Math,
	TasksOfBlock16Math,
	TasksOfBlock17Math,
	TasksOfBlock18Math,
	TasksOfBlock19Math,
	TasksOfBlock20Math,
	TasksOfBlock21Math,
}