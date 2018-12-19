package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/CookiesChen/AI/SA/hill_climbing"
)

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	filePath := "tsp/d198.tsp"
	xs, ys := getData(filePath)
	hill_climbing.Exec(xs, ys)
	// sa.Exec(xs, ys)
	<-c
}

// 读取文件
func getData(filePath string) (xs []float64, ys []float64) {
	dataStr := `1 0.00000e+00 0.00000e+00
2 5.51200e+02 9.96400e+02
3 6.27400e+02 9.96400e+02
4 7.03600e+02 9.96400e+02
5 7.03600e+02 1.04720e+03
6 6.27400e+02 1.04720e+03
7 5.51200e+02 1.04720e+03
8 8.81400e+02 1.35200e+03
9 9.32200e+02 1.35200e+03
10 9.57600e+02 1.35200e+03
11 9.83000e+02 1.35200e+03
12 1.00840e+03 1.35200e+03
13 1.03380e+03 1.35200e+03
14 1.31320e+03 1.12340e+03
15 1.28780e+03 1.09800e+03
16 1.28780e+03 9.96400e+02
17 1.31320e+03 9.96400e+02
18 1.46560e+03 9.96400e+02
19 1.51640e+03 9.96400e+02
20 1.59260e+03 9.96400e+02
21 1.59260e+03 1.09800e+03
22 1.51640e+03 1.09800e+03
23 1.46560e+03 1.09800e+03
24 1.56720e+03 1.12340e+03
25 1.59260e+03 1.14880e+03
26 1.56720e+03 1.17420e+03
27 1.54180e+03 1.17420e+03
28 1.49100e+03 1.17420e+03
29 1.44020e+03 1.17420e+03
30 1.46560e+03 1.19960e+03
31 1.41480e+03 1.22500e+03
32 1.44020e+03 1.22500e+03
33 1.49100e+03 1.22500e+03
34 1.51640e+03 1.22500e+03
35 1.59260e+03 1.25040e+03
36 1.46560e+03 1.25040e+03
37 1.44020e+03 1.25040e+03
38 1.38940e+03 1.25040e+03
39 1.54180e+03 1.27580e+03
40 1.16080e+03 1.12340e+03
41 1.16080e+03 1.22500e+03
42 1.26240e+03 1.30120e+03
43 1.28780e+03 1.30120e+03
44 1.33860e+03 1.30120e+03
45 1.41480e+03 1.30120e+03
46 1.49100e+03 1.30120e+03
47 1.54180e+03 1.30120e+03
48 1.64340e+03 1.30120e+03
49 1.66880e+03 1.32660e+03
50 1.61800e+03 1.32660e+03
51 1.56720e+03 1.32660e+03
52 1.51640e+03 1.32660e+03
53 1.46560e+03 1.32660e+03
54 1.41480e+03 1.32660e+03
55 1.33860e+03 1.32660e+03
56 1.31320e+03 1.32660e+03
57 1.23700e+03 1.32660e+03
58 1.23700e+03 1.35200e+03
59 1.31320e+03 1.35200e+03
60 1.33860e+03 1.35200e+03
61 1.41480e+03 1.35200e+03
62 1.46560e+03 1.35200e+03
63 1.69420e+03 1.35200e+03
64 1.61800e+03 1.37740e+03
65 1.51640e+03 1.37740e+03
66 1.41480e+03 1.37740e+03
67 1.33860e+03 1.37740e+03
68 1.31320e+03 1.37740e+03
69 1.23700e+03 1.37740e+03
70 1.35770e+03 1.39010e+03
71 1.23700e+03 1.40280e+03
72 1.31320e+03 1.40280e+03
73 1.33860e+03 1.40280e+03
74 1.41480e+03 1.40280e+03
75 1.46560e+03 1.40280e+03
76 1.56720e+03 1.40280e+03
77 1.59260e+03 1.40280e+03
78 1.61800e+03 1.40280e+03
79 1.69420e+03 1.42820e+03
80 1.66880e+03 1.42820e+03
81 1.54180e+03 1.42820e+03
82 1.44020e+03 1.42820e+03
83 1.41480e+03 1.42820e+03
84 1.33860e+03 1.42820e+03
85 1.26240e+03 1.42820e+03
86 1.23700e+03 1.45360e+03
87 1.33860e+03 1.45360e+03
88 1.41480e+03 1.45360e+03
89 1.46560e+03 1.45360e+03
90 1.49100e+03 1.45360e+03
91 1.66880e+03 1.45360e+03
92 1.69420e+03 1.45360e+03
93 1.61800e+03 1.47900e+03
94 1.59260e+03 1.47900e+03
95 1.56720e+03 1.47900e+03
96 1.51640e+03 1.47900e+03
97 1.44020e+03 1.47900e+03
98 1.41480e+03 1.47900e+03
99 1.33860e+03 1.47900e+03
100 1.26240e+03 1.47900e+03
101 1.59260e+03 1.50440e+03
102 1.61800e+03 1.50440e+03
103 1.66880e+03 1.52980e+03
104 1.69420e+03 1.52980e+03
105 1.74500e+03 1.52980e+03
106 1.82120e+03 1.52980e+03
107 1.84660e+03 1.52980e+03
108 1.94820e+03 1.52980e+03
109 1.92280e+03 1.55520e+03
110 1.89740e+03 1.55520e+03
111 1.87200e+03 1.55520e+03
112 1.82120e+03 1.55520e+03
113 1.79580e+03 1.55520e+03
114 1.77040e+03 1.55520e+03
115 1.77040e+03 1.65680e+03
116 1.79580e+03 1.65680e+03
117 1.82120e+03 1.65680e+03
118 1.87200e+03 1.65680e+03
119 1.89740e+03 1.65680e+03
120 1.92280e+03 1.65680e+03
121 1.80850e+03 1.69490e+03
122 1.75770e+03 1.69490e+03
123 1.88470e+03 1.73300e+03
124 1.99900e+03 1.73300e+03
125 2.07520e+03 1.73300e+03
126 2.11330e+03 1.73300e+03
127 2.17680e+03 1.73300e+03
128 2.23650e+03 1.73300e+03
129 2.17680e+03 1.78380e+03
130 2.12600e+03 1.78380e+03
131 2.10060e+03 1.78380e+03
132 2.10060e+03 1.80920e+03
133 2.12600e+03 1.80920e+03
134 2.10060e+03 1.83460e+03
135 2.12600e+03 1.83460e+03
136 2.15140e+03 1.83460e+03
137 2.23650e+03 1.84730e+03
138 1.99900e+03 1.84730e+03
139 1.88470e+03 1.84730e+03
140 2.10060e+03 1.86000e+03
141 2.12600e+03 1.86000e+03
142 2.10060e+03 1.88540e+03
143 2.12600e+03 1.88540e+03
144 2.17680e+03 1.88540e+03
145 2.15140e+03 1.91080e+03
146 2.12600e+03 1.91080e+03
147 2.10060e+03 1.91080e+03
148 2.10060e+03 1.93620e+03
149 2.12600e+03 1.93620e+03
150 2.17680e+03 1.93620e+03
151 2.22760e+03 1.93620e+03
152 2.12600e+03 1.96160e+03
153 2.10060e+03 1.96160e+03
154 1.79580e+03 1.98700e+03
155 1.82120e+03 1.98700e+03
156 1.84660e+03 1.98700e+03
157 1.87200e+03 1.98700e+03
158 1.89740e+03 1.98700e+03
159 1.94820e+03 1.98700e+03
160 2.05620e+03 1.98700e+03
161 2.10060e+03 1.98700e+03
162 2.12600e+03 1.98700e+03
163 2.25300e+03 1.98700e+03
164 2.30380e+03 1.98700e+03
165 2.38000e+03 1.98700e+03
166 2.40540e+03 1.98700e+03
167 2.02440e+03 1.40280e+03
168 2.15140e+03 1.40280e+03
169 2.07520e+03 1.70760e+03
170 2.17680e+03 1.70760e+03
171 2.35080e+03 1.73300e+03
172 2.35080e+03 1.84730e+03
173 3.65210e+03 1.01030e+03
174 3.72570e+03 1.01030e+03
175 3.72570e+03 1.08650e+03
176 3.65210e+03 1.08650e+03
177 3.72620e+03 1.14880e+03
178 3.80240e+03 1.14880e+03
179 3.85320e+03 1.14880e+03
180 3.80240e+03 1.17420e+03
181 3.70080e+03 1.17420e+03
182 3.60560e+03 1.19960e+03
183 3.70080e+03 1.19960e+03
184 3.80240e+03 1.19960e+03
185 3.85320e+03 1.19960e+03
186 4.02830e+03 1.31030e+03
187 3.95210e+03 1.31030e+03
188 3.72830e+03 1.31030e+03
189 3.65210e+03 1.31030e+03
190 3.65210e+03 1.38650e+03
191 3.72830e+03 1.38650e+03
192 3.95210e+03 1.38650e+03
193 4.02830e+03 1.38650e+03
194 3.85320e+03 1.12340e+03
195 3.95210e+03 1.08650e+03
196 4.02830e+03 1.08650e+03
197 4.02830e+03 1.01030e+03
198 3.95210e+03 1.01030e+03
EOF
`
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return
	// }
	// defer file.Close()
	rd := bufio.NewReader(strings.NewReader(dataStr))
	for {
		line, _, _ := rd.ReadLine()
		if strings.Contains(string(line), "EOF") {
			break
		}
		data := string(line)
		lineData := strings.Split(data, " ")
		coordinateX, _ := strconv.ParseFloat(lineData[1], 64)
		coordinateY, _ := strconv.ParseFloat(lineData[2], 64)
		xs = append(xs, coordinateX)
		ys = append(ys, coordinateY)
	}
	return xs, ys
}
