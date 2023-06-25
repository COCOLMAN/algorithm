package LongestIncreasingSubsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "default",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "default",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "default",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7},
			},
			want: 7,
		},
		{
			name: "default",
			args: args{
				nums: []int{4951, 4865, 6665, 1353, 5818, 9505, 3600, 3942, 9696, 1234, 6278, 9692, 2001, 433, 1816, 9563, 8864, 4860, 2390, 5756, 5590, 1178, 8350, 2095, 6923, 9567, 5183, 5168, 3797, 197, 1871, 7678, 1402, 4222, 522, 5460, 7163, 2287, 97, 8952, 747, 1991, 8850, 6307, 368, 8317, 1059, 5644, 1642, 2334, 2104, 1658, 7570, 9912, 9184, 1846, 5640, 8417, 120, 7196, 773, 6460, 3996, 8563, 2220, 7659, 7976, 7364, 8890, 8327, 8987, 9799, 2627, 3656, 4992, 7352, 916, 3366, 855, 2313, 439, 6747, 2470, 1218, 4868, 4863, 6643, 2158, 6210, 7406, 745, 3871, 6366, 6891, 3735, 7762, 3972, 443, 413, 4453, 2878, 9066, 7380, 721, 9137, 9051, 4456, 7663, 5568, 6755, 6239, 1858, 1904, 9025, 1102, 9513, 2772, 6954, 924, 358, 5358, 2733, 7998, 9128, 6653, 5162, 5436, 2235, 1477, 655, 7395, 6038, 6805, 903, 7780, 8356, 3261, 7625, 4081, 2372, 6135, 2648, 8753, 8886, 1016, 7151, 5587, 8494, 2948, 743, 9539, 3564, 369, 1766, 4449, 8584, 1080, 9755, 7636, 2993, 6459, 3680, 5154, 2169, 8925, 6012, 6982, 319, 8114, 8655, 2399, 4736, 692, 1070, 4993, 2588, 7484, 7444, 8767, 3197, 7146, 9896, 7479, 1458, 6694, 6190, 1129, 3474, 2068, 5682, 8364, 8535, 2958, 9072, 3131, 4871, 6338, 4234, 4077, 1676, 6251, 218, 1170, 4803, 5636, 7856, 8466, 6528, 2223, 2236, 6986, 1538, 7180, 4167, 7272, 478, 384, 6027, 7166, 8733, 6062, 5239, 538, 8227, 460, 2490, 3907, 9759, 2669, 678, 1515, 4891, 2795, 4171, 4761, 7116, 2586, 6709, 1059, 9737, 1874, 7592, 6420, 7496, 4575, 9140, 4414, 4910, 3633, 3656, 3600, 2017, 9923, 539, 6746, 9104, 7925, 4, 8277, 1484, 354, 1574, 7751, 1270, 8914, 1143, 2992, 2583, 4487, 5280, 9733, 2205, 8492, 5565, 1000, 6377, 2456, 902, 2043, 8743, 8250, 3894, 1081, 8993, 7375, 2196, 2605, 9065, 4401, 4804, 5632, 7202, 8361, 900, 5674, 9788, 274, 3841, 8464, 171, 1519, 4454, 6049, 1710, 3861, 4517, 7398, 5401, 4299, 5242, 5364, 2372, 6407, 1240, 4507, 9978, 4291, 4002, 4387, 4965, 8783, 5607, 2213, 1533, 5395, 3685, 275, 4958, 1275, 1606, 7778, 3067, 5984, 763, 1090, 8562, 6971, 5434, 9233, 9142, 9262, 6313, 2668, 6336, 9215, 2923, 7067, 927, 7387, 6531, 7016, 4294, 7355, 8912, 445, 3703, 8510, 9751, 2052, 9679, 4179, 4534, 2361, 4583, 9981, 6533, 3485, 2854, 3302, 4584, 5928, 3410, 5569, 7645, 2802, 3954, 4744, 132, 2532, 6994, 7582, 3862, 8603, 5458, 3181, 6832, 4191, 2286, 554, 6633, 692, 1600, 9120, 2077, 4207, 8496, 9810, 9397, 3447, 4551, 288, 8210, 4342, 1629, 2894, 3024, 4063, 1455, 9086, 290, 3122, 1218, 7897, 6479, 9210, 6653, 7557, 7238, 2740, 7456, 154, 989, 4583, 7973, 1889, 7464, 6262, 5891, 9668, 246, 7154, 7742, 2114, 1269, 9336, 3106, 9540, 359, 6079, 1385, 755, 4331, 1340, 5116, 8957, 5366, 4638, 2830, 3906, 9867, 5118, 8920, 8141, 8212, 2819, 2341, 2103, 8918, 1125, 6429, 3514, 5644, 854, 6210, 9965, 4863, 858, 5524, 6901, 5045, 3742, 4367, 6722, 3190, 6357, 8701, 5104, 4519, 8726, 5457, 9742, 6017, 1729, 6756, 4598, 8785, 5768, 8801, 600, 1531, 5546, 692, 2956, 9326, 8765, 5044, 5015, 981, 9549, 2241, 9067, 2653, 5931, 9957, 8793, 6724, 5633, 6918, 6020, 7510, 9365, 9436, 5687, 531, 7041, 9816, 6136, 6827, 2987, 1696, 5418, 3409, 3796, 2899, 5963, 3177, 834, 2713, 2464, 5514, 8827, 8401, 4082, 4921, 9004, 6007, 3268, 2390, 9581, 6518, 1702, 495, 3187, 7544, 3647, 7782, 9232, 5196, 846, 3916, 2781, 9824, 9473, 7144, 1743, 9727, 8569, 1584, 4559, 5966, 9814, 8195, 5095, 4812, 8000, 6105, 8203, 8200, 7866, 5622, 4634, 236, 7843, 7421, 8920, 6278, 1717, 9032, 9134, 1850, 9967, 563, 8595, 6344, 6146, 6781, 1709, 5196, 500, 161, 8940, 2839, 3261, 6542, 7801, 5718, 5247, 2630, 750, 7689, 1133, 5991, 4081, 5705, 1713, 4684, 7517, 9963, 9118, 2647, 5699, 9531, 10, 5728, 7024, 9541, 422, 7868, 569, 6562, 4098, 8462, 6134, 1020, 8121, 4052, 1819, 5546, 2378, 7434, 703, 199, 5337, 1024, 3479, 8958, 7876, 2117, 5139, 9760, 968, 5810, 3036, 7157, 2488, 5374, 4442, 1235, 3044, 5382, 1165, 1653, 2699, 6942, 2352, 7032, 9663, 9682, 7031, 8895, 1179, 6585, 4976, 5002, 2605, 8899, 5340, 3919, 8444, 4906, 6995, 6766, 478, 6312, 7264, 9120, 5402, 1013, 1927, 8058, 7872, 4002, 7033, 3359, 5298, 6788, 6403, 3306, 7662, 6941, 480, 298, 6412, 2266, 1893, 5452, 8012, 8098, 3251, 9183, 6711, 44, 8842, 6511, 8994, 1701, 4174, 4732, 9285, 4692, 6838, 8014, 8408, 4655, 5593, 5523, 8284, 3072, 3798, 2823, 2699, 802, 2918, 7307, 7766, 4104, 4335, 166, 2879, 9382, 6316, 9227, 7518, 9298, 7906, 4044, 7918, 62, 6166, 3180, 6523, 8316, 625, 9559, 3061, 4091, 1762, 6343, 3273, 9498, 8130, 5177, 1736, 9509, 85, 9798, 6758, 5323, 1128, 3933, 4814, 583, 2073, 835, 7473, 4480, 8444, 5585, 9041, 9483, 9589, 4096, 2265, 6999, 4791, 2558, 9797, 3241, 6275, 4099, 7051, 1468, 1397, 3295, 1757, 3296, 7959, 3875, 8571, 7253, 5615, 766, 499, 9045, 1881, 4123, 2168, 2201, 6635, 8954, 106, 1055, 5799, 8425, 4643, 482, 48, 5180, 3034, 1812, 8943, 3916, 5481, 2553, 3022, 215, 1350, 4599, 7818, 8197, 8666, 9823, 4818, 3182, 8173, 5190, 7400, 7264, 5614, 8953, 7278, 5950, 9925, 7034, 8020, 4152, 8664, 5580, 9741, 1041, 3407, 8996, 7646, 2123, 7494, 7435, 9217, 7781, 8868, 1259, 4092, 1826, 5770, 1049, 9972, 346, 1462, 4546, 9731, 2593, 864, 9881, 2821, 4348, 5130, 6172, 3688, 4603, 7278, 4821, 8198, 1618, 2147, 3782, 5811, 3539, 4448, 926, 8281, 3715, 3283, 9800, 832, 1651, 2202, 8431, 9719, 9140, 4863, 2661, 347, 3753, 1146, 3502, 5788, 9487, 4694, 3337, 131, 3519, 1637, 5077, 533, 4396, 3532, 5994, 7877, 169, 251, 3143, 6734, 8535, 3021, 7364, 7322, 1501, 4846, 2031, 3411, 9081, 4653, 4502, 5512, 565, 4961, 2030, 7485, 8952, 7991, 1065, 2571, 2061, 3972, 1809, 5862, 1534, 4125, 7291, 8375, 3258, 5098, 3799, 110, 9449, 9797, 1580, 1646, 4768, 241, 3766, 7194, 6460, 4087, 7332, 4932, 3302, 4566, 6719, 9219, 4423, 8074, 2400, 8551, 2849, 809, 12, 7201, 8678, 8434, 9141, 7587, 9210, 5035, 1535, 2530, 5698, 7343, 2109, 9859, 2838, 2897, 6012, 7171, 9540, 9231, 5562, 7953, 5542, 3006, 3668, 4503, 2499, 5155, 2795, 9377, 7439, 5446, 6426, 7818, 1099, 7332, 6248, 2217, 5759, 9176, 2425, 1465, 3693, 7588, 7556, 5396, 4002, 6170, 4824, 3250, 8564, 7212, 4794, 1638, 3126, 170, 5759, 8522, 3555, 6236, 7680, 9160, 3681, 1464, 1607, 2149, 2905, 2324, 1471, 3385, 3569, 7252, 7425, 3992, 4804, 5762, 458, 8780, 5969, 8753, 5576, 1903, 9518, 9216, 1272, 3061, 1806, 7117, 4724, 5136, 5701, 1322, 3584, 3764, 5745, 354, 6302, 2120, 4193, 4481, 8288, 4791, 846, 302, 4002, 5804, 7136, 2670, 8605, 6851, 6253, 9496, 8254, 7987, 221, 4628, 6450, 5842, 738, 8497, 2643, 2258, 5249, 5133, 2478, 7045, 6393, 3753, 8124, 2219, 3412, 2531, 3363, 9282, 8947, 1494, 5751, 5441, 3574, 6417, 2785, 7589, 5450, 1390, 4331, 1891, 3381, 3913, 626, 8102, 2968, 6270, 1098, 7068, 4894, 5237, 11, 9194, 8887, 4054, 3240, 6857, 1942, 7931, 6438, 9518, 2427, 2586, 9314, 5166, 7675, 4052, 5484, 9321, 9458, 7969, 1696, 5103, 3143, 9026, 198, 7314, 1490, 7876, 7318, 1079, 8991, 9695, 1246, 3217, 1244, 2037, 7518, 231, 2748, 3230, 2335, 1885, 4373, 4396, 8713, 2671, 6617, 5264, 4065, 1359, 8454, 9990, 3830, 430, 6890, 2652, 4669, 7161, 9789, 6401, 4590, 7239, 5678, 1186, 4625, 660, 1691, 8934, 2133, 9221, 7134, 2470, 5087, 2170, 1376, 6025, 8564, 8198, 6629, 8645, 1334, 7013, 2375, 2353, 8036, 7300, 9890, 4864, 3375, 6392, 5743, 2810, 8734, 4750, 3776, 6209, 2305, 234, 2814, 8487, 8124, 2794, 2219, 7879, 7487, 4838, 8165, 2327, 530, 5921, 1537, 3123, 4197, 2744, 9108, 1855, 6789, 7437, 5718, 5859, 8024, 9111, 4786, 6293, 267, 7165, 459, 3015, 2588, 892, 6666, 2078, 5175, 7808, 6973, 4534, 8844, 5931, 7684, 7624, 5209, 8688, 7561, 4632, 5601, 6978, 8322, 378, 2161, 2066, 9896, 4078, 4952, 6089, 209, 9162, 7553, 3224, 8854, 6033, 5221, 7598, 2070, 3338, 2022, 320, 6985, 3806, 4652, 4331, 95, 9561, 9070, 6793, 3616, 7469, 8389, 9822, 6055, 1647, 5398, 7612, 7602, 7609, 4903, 4131, 4640, 514, 4477, 9593, 6161, 8330, 7289, 6566, 2109, 2161, 2517, 5458, 1262, 2745, 5181, 6086, 4908, 3904, 4301, 2877, 1039, 4706, 6889, 7439, 3312, 6455, 7249, 8713, 6804, 798, 9195, 4392, 8270, 6667, 5030, 3932, 4690, 3540, 3434, 1980, 3411, 3804, 7672, 9315, 8463, 2584, 1444, 1678, 177, 4153, 7311, 1021, 64, 6727, 6619, 570, 815, 6315, 1702, 1964, 3756, 629, 9499, 8305, 5699, 6096, 3578, 132, 1999, 72, 9450, 3009, 2621, 8183, 2453, 1191, 9331, 3162, 5065, 215, 6789, 6902, 6281, 9788, 2862, 1444, 5256, 9145, 3528, 6781, 6550, 6177, 2491, 3692, 1287, 3737, 1133, 4378, 4737, 8732, 372, 906, 8726, 9452, 731, 11, 6894, 990, 9997, 9908, 5097, 6980, 7754, 4256, 9270, 7954, 703, 5860, 1338, 1347, 294, 1687, 7128, 2195, 1839, 6795, 9152, 9646, 1730, 9269, 1952, 2428, 8038, 8920, 5974, 1285, 1655, 9234, 9286, 3997, 6693, 3368, 3115, 5068, 4796, 1046, 9682, 1077, 3168, 8521, 4961, 5365, 828, 2567, 230, 2861, 9951, 6810, 2426, 3396, 6758, 4427, 3847, 8466, 6436, 951, 9338, 10, 4805, 7022, 6146, 2759, 8092, 5644, 8969, 9656, 6203, 5753, 5058, 102, 6312, 1095, 8578, 8788, 4596, 3824, 2500, 3044, 7949, 5623, 9120, 7442, 5941, 8407, 1406, 8796, 8463, 9115, 6373, 4623, 7907, 6148, 7855, 557, 8961, 2831, 290, 6566, 9831, 3725, 412, 6902, 2919, 8713, 1352, 7025, 8771, 8854, 4287, 1547, 3789, 4303, 6496, 7544, 4166, 2553, 9336, 2111, 6271, 1546, 6276, 5123, 2477, 6769, 2615, 9722, 8047, 7915, 2926, 9099, 2889, 9286, 3915, 7806, 3838, 4654, 535, 9947, 8741, 7146, 7753, 4912, 2049, 4212, 6581, 2009, 3004, 7262, 6156, 1926, 4852, 2624, 2245, 4716, 3327, 8592, 1363, 9532, 5811, 229, 8398, 9291, 4821, 7918, 7909, 9263, 9261, 6223, 3874, 6538, 6038, 6110, 2938, 644, 1446, 7132, 107, 8062, 9430, 3431, 7835, 7979, 4023, 3567, 7433, 433, 8551, 3413, 5211, 8930, 9152, 5860, 9901, 8718, 4651, 8088, 3439, 5366, 4428, 4579, 6776, 9119, 4501, 2291, 2979, 5393, 4486, 9664, 5714, 8264, 5238, 4942, 2265, 1788, 480, 1538, 1416, 181, 4705, 7022, 5768, 1594, 2088, 8722, 3665, 2271, 7852, 8093, 7063, 2564, 5001, 8029, 4827, 3817, 8838, 8699, 184, 9866, 7644, 7231, 8126, 4526, 7323, 5725, 1634, 2914, 1260, 4866, 468, 5180, 9066, 7272, 5888, 7093, 6522, 1947, 2724, 3235, 2995, 3977, 7749, 1728, 394, 3837, 3332, 4411, 2656, 6542, 5755, 9993, 1687, 3158, 1928, 5716, 142, 2730, 3758, 7648, 7044, 9045, 68, 1507, 8372, 1017, 2102, 213, 7879, 2368, 3444, 1537, 45, 903, 1662, 2908, 9355, 6335, 697, 8885, 9706, 7072, 6752, 4215, 1498, 4495, 7597, 6228, 316, 4679, 1482, 6203, 2954, 2863, 4687, 2051, 6046, 1927, 3550, 8579, 7456, 2319, 7818, 8403, 207, 8034, 6906, 5394, 8464, 5420, 2525, 473, 9869, 2444, 7429, 636, 7336, 3433, 3610, 7759, 3515, 253, 2490, 2564, 6674, 5889, 7431, 9929, 2865, 6813, 6411, 897, 4612, 1987, 9254, 4910, 3778, 5854, 9456, 1644, 5581, 863, 3128, 1114, 3118, 5698, 6539, 2244, 6502, 4049, 6468, 6077, 5691, 9632, 7318, 6354, 797, 5764, 7075, 2280, 2817, 2957, 2897, 1335, 2303, 9978, 8538, 532, 3974, 6663, 3178, 8819, 2041, 4632, 7698, 7224, 1587, 6132, 8931, 4461, 8328, 6662, 5692, 8993, 5997, 4329, 7653, 3683, 5350, 200, 4355, 6688, 7833, 3174, 3403, 7147, 7113, 2349, 3462, 5142, 3251, 4676, 453, 473, 3881, 909, 6338, 8497, 3008, 7287, 3263, 8226, 1085, 3829, 7406, 193, 4197, 5800, 207, 8113, 9298, 8779, 4680, 5143, 8758, 1091, 3536, 7182, 6600, 652, 1565, 3675, 6304, 8378, 5969, 3326, 6583, 4652, 4937, 5058, 3619, 3750, 9893, 9429, 2222, 510, 4926, 4847, 4645, 8795, 2917, 2230, 8339, 480, 7444, 8765, 1230, 5535, 6773, 2627, 9189, 8730, 7725, 1173, 3286, 3628, 4837, 4470, 8688, 5170, 67, 8031, 9090, 2179, 4860, 9784, 7746, 7719, 5308, 8275, 4173, 7705, 187, 218, 776, 6911, 9427, 6753, 6790, 3342, 4908, 5616, 2680, 680, 1763, 6432, 2096, 1749, 9477, 3576, 9002, 4518, 9241, 1045, 7317, 1327, 1018, 8548, 9712, 8754, 8289, 3514, 8448, 7238, 3129, 9063, 2103, 2025, 1887, 7315, 1426, 6516, 9383, 748, 2358, 437, 4104, 7798, 5563, 3874, 1853, 8858, 9380, 4275, 5116, 3965, 9432, 9200, 638, 4287, 5636, 872, 7816, 4329, 3280, 3495, 10, 276, 5927, 2380, 9143, 172, 1079, 4610, 9015, 4142, 3186, 5149, 3886, 2408, 228, 3196, 8344, 1459, 4337},
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSearchIdx(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "defualt",
			args: args{
				target: 3,
				nums:   []int{1, 3},
			},
			want: 1,
		},
		{
			name: "defualt",
			args: args{
				target: 3,
				nums:   []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "defualt",
			args: args{
				target: 3,
				nums:   []int{1, 2, 4},
			},
			want: 2,
		},
		{
			name: "defualt",
			args: args{
				target: 3,
				nums:   []int{3, 4},
			},
			want: 0,
		},
		{
			name: "defualt",
			args: args{
				target: 3,
				nums:   []int{3, 4, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BSearchIdx(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("BSearchIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}
