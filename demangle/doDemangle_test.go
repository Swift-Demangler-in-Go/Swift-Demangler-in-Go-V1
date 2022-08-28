package demangle

import (
	"testing"
)

func TestDemangler(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"_TFC4Pack5class4FuncFT_Si", //没有参数
			"Pack.class.Func() -> Swift.Int",
		},
		{
			"_TFC4Pack5class4FuncFT_Sb",
			"Pack.class.Func() -> Swift.Bool",
		},
		{
			"_TFC4Pack5class4FuncFT_S0_",
			"Pack.class.Func() -> Pack.class",
		},
		{
			"_TFC4Pack5class4FuncFT_SS",
			"Pack.class.Func() -> Swift.String",
		},
		{
			"_TFCC4Pack6classA6classB4FuncFT_S1_",
			"Pack.classA.classB.Func() -> Pack.classA.classB",
		},
		{
			"_TFC4Pack5class4FuncFSiT_", //没有返回值
			"Pack.class.Func(Swift.Int) -> ()",
		},
		{
			"_TFC4Pack5class4FuncFS0_T_",
			"Pack.class.Func(Pack.class) -> ()",
		},
		{
			"_TFCC4pack6classA6classB4FuncFFFSiSSFSbS0_T_",
			"pack.classA.classB.Func(((Swift.Int) -> Swift.String) -> (Swift.Bool) -> pack.classA) -> ()",
		},
		{
			"_TFC4Pack5class4FuncFSiSi", //有参数和返回值
			"Pack.class.Func(Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4Pack5class4FuncFSSSb",
			"Pack.class.Func(Swift.String) -> Swift.Bool",
		},
		{
			"_TFCC4pack6classA6classB4FuncFSiS1_",
			"pack.classA.classB.Func(Swift.Int) -> pack.classA.classB",
		},
		{
			"_TFC4Pack5class4FuncFTSiSb_SS", //元组参数
			"Pack.class.Func(Swift.Int, Swift.Bool) -> Swift.String",
		},
		{"_TFC4Pack5class4FuncFTSSSi_Sb",
			"Pack.class.Func(Swift.String, Swift.Int) -> Swift.Bool",
		},
		{
			"_TFC4Pack5class4FuncFSiTSiSb_", //元组返回值
			"Pack.class.Func(Swift.Int) -> (Swift.Int, Swift.Bool)",
		},
		{
			"_TFC4Pack5class4FuncFFFFSiSiSiSiTSiSb_",
			"Pack.class.Func((((Swift.Int) -> Swift.Int) -> Swift.Int) -> Swift.Int) -> (Swift.Int, Swift.Bool)",
		},
		{
			"_TFC4Pack5class4FuncFS0_TSiSS_",
			"Pack.class.Func(Pack.class) -> (Swift.Int, Swift.String)",
		},
		{
			"_TFC4Pack5class4FuncFTSiSb_TSSSi_", //元组参数和元组返回值
			"Pack.class.Func(Swift.Int, Swift.Bool) -> (Swift.String, Swift.Int)",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTS0_S1__TSSSi_",
			"pack.classA.classB.Func(pack.classA, pack.classA.classB) -> (Swift.String, Swift.Int)",
		},
		{
			"_TF4pack4funcFtSiSS_Si", //可变长度元组
			"pack.func(Swift.Int, Swift.String...) -> Swift.Int",
		},
		{
			"_TFC4pack5class4funcFtSiSa_Si",
			"pack.class.func(Swift.Int, Swift.Array...) -> Swift.Int",
		},
		{
			"_TFC4Pack5classg4FuncFTt1aSiSf__SS",
			"Pack.class.Func.getter : ((a: Swift.Int, Swift.Float...)) -> Swift.String",
		},
		{
			"_TFC4Pack5class4FuncFTt1aSiSf__tfSSSS_", //返回值为可变长度元组
			"Pack.class.Func((a: Swift.Int, Swift.Float...)) -> ((Swift.String) -> Swift.String...)",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTSiSfSb_tfSbSS_",
			"pack.classA.classB.Func(Swift.Int, Swift.Float, Swift.Bool) -> ((Swift.Bool) -> Swift.String...)",
		},
		{
			"_TFC4pack5class4funcFT1fSiSb_FtSi_Si",
			"pack.class.func(f: Swift.Int, Swift.Bool) -> (Swift.Int...) -> Swift.Int",
		},
		{
			"_TFC4pack5class4funcFT1fSiSb_tSi_",
			"pack.class.func(f: Swift.Int, Swift.Bool) -> (Swift.Int...)",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTtSiSSSiSb__tfSSSS_",
			"pack.classA.classB.Func((Swift.Int, Swift.String, Swift.Int, Swift.Bool...)) -> ((Swift.String) -> Swift.String...)",
		},
		{
			"_TFC4Pack5class4FuncFFSiSbSS", //参数是闭包
			"Pack.class.Func((Swift.Int) -> Swift.Bool) -> Swift.String",
		},
		{
			"_TFC4Pack5class4FuncFFSSS0_Si",
			"Pack.class.Func((Swift.String) -> Pack.class) -> Swift.Int",
		},
		{
			"_TFC4Pack5class4FuncFTSiS0__FSbSS",
			"Pack.class.Func(Swift.Int, Pack.class) -> (Swift.Bool) -> Swift.String",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTSiS1__FFSiSiS0_",
			"pack.classA.classB.Func(Swift.Int, pack.classA.classB) -> ((Swift.Int) -> Swift.Int) -> pack.classA",
		},
		{
			"_TFC4Pack5class4FuncFSiFSbSS", //返回值是闭包
			"Pack.class.Func(Swift.Int) -> (Swift.Bool) -> Swift.String",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTS1__FFSiSiSS",
			"pack.classA.classB.Func(pack.classA.classB) -> ((Swift.Int) -> Swift.Int) -> Swift.String",
		},
		{
			"_TFC4pack5class4FuncFFSiSiFSiSi", //参数和返回值都是闭包
			"pack.class.Func((Swift.Int) -> Swift.Int) -> (Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4Pack5class4FuncFTSiSb_FFSbSSS0_",
			"Pack.class.Func(Swift.Int, Swift.Bool) -> ((Swift.Bool) -> Swift.String) -> Pack.class",
		},
		{
			"_TFC4pack5class4FuncFFFSiSSSSSi", //参数是闭包且闭包的参数是闭包
			"pack.class.Func(((Swift.Int) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFFSiSbSSSbS0_",
			"pack.class.Func((((Swift.Int) -> Swift.Bool) -> Swift.String) -> Swift.Bool) -> pack.class",
		},
		{
			"_TFC4pack5class4FuncFFSiFSbSSSi", //参数是闭包且闭包的返回值是闭包
			"pack.class.Func((Swift.Int) -> (Swift.Bool) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFSiSSFSbSSSi", //参数是闭包且闭包的参数、返回值是闭包
			"pack.class.Func(((Swift.Int) -> Swift.String) -> (Swift.Bool) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFCC4pack6classA6classB4FuncFFFS0_SbSSS0_",
			"pack.classA.classB.Func(((pack.classA) -> Swift.Bool) -> Swift.String) -> pack.classA",
		},
		{
			"_TFCC4pack6classA6classB4FuncFFFSiSSFSbS0_Si",
			"pack.classA.classB.Func(((Swift.Int) -> Swift.String) -> (Swift.Bool) -> pack.classA) -> Swift.Int",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTFFSiSbSf_S1_",
			"pack.classA.classB.Func(((Swift.Int) -> Swift.Bool) -> Swift.Float) -> pack.classA.classB",
		},
		{
			"_TFCC4pack6classA6classB4FuncFTFFSiSbSf_ffSfSbSS",
			"pack.classA.classB.Func(((Swift.Int) -> Swift.Bool) -> Swift.Float) -> ((Swift.Float) -> Swift.Bool) -> Swift.String",
		},
		{
			"_TFC4test5class4funcfT1aSi1bSf_Sb", //外部参数类型
			"test.class.func(a: Swift.Int, b: Swift.Float) -> Swift.Bool",
		},
		{
			"_TFCC4pack6classA6classB4FuncFT1aSi1bSb1cSS_S1_",
			"pack.classA.classB.Func(a: Swift.Int, b: Swift.Bool, c: Swift.String) -> pack.classA.classB",
		},
		{
			"_TFC4test5class4funcfT1aSi1bSf_T1dSb1eSi_",
			"test.class.func(a: Swift.Int, b: Swift.Float) -> (d: Swift.Bool, e: Swift.Int)",
		},
		{
			"_TFC4pack5class4funcFT1aSi1bSf1cSS_TSb1dfSiSf_", //内部参数类型
			"pack.class.func(a: Swift.Int, b: Swift.Float, c: Swift.String) -> (Swift.Bool, d: (Swift.Int) -> Swift.Float)",
		},
		{
			"_TFCC4pack6classA6classB4FuncFT1aSiSi_T1efS1_SSSS_",
			"pack.classA.classB.Func(a: Swift.Int, Swift.Int) -> (e: (pack.classA.classB) -> Swift.String, Swift.String)",
		},
		{
			"_TFC4test5class4funcfT1aSi1bSf_T1QSSFSSSS_", //外部参数
			"test.class.func(a: Swift.Int, b: Swift.Float) -> (Q: Swift.String, (Swift.String) -> Swift.String)",
		},
		{
			"_TF1Q1QFTFSiSbSf_TSi1xFSbSf_",
			"Q.Q((Swift.Int) -> Swift.Bool, Swift.Float) -> (Swift.Int, x: (Swift.Bool) -> Swift.Float)",
		},
		{
			"_TFC4Pack5class4FuncFTFSiSbSf_T1eFSbSf_",
			"Pack.class.Func((Swift.Int) -> Swift.Bool, Swift.Float) -> (e: (Swift.Bool) -> Swift.Float)",
		},
		{
			"_TFC4pack5class4FuncFFS0_SSSb",
			"pack.class.Func((pack.class) -> Swift.String) -> Swift.Bool",
		},
		{
			"_TFC4pack5classW4funcFSiSi", //（出现在func前的多余的g，s 等输出在func后面的操作）
			"pack.class.func.didset : (Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4pack5classg4funcFT_fSiSS",
			"pack.class.func.getter : () -> (Swift.Int) -> Swift.String",
		},
		{
			"_TFC4pack5classao4funcFSbSS",
			"pack.class.func.nativeOwningMutableAddressor : (Swift.Bool) -> Swift.String",
		},
		{
			"_TFC4pack5classlO4funcFT1aSi1bSi_T1cSi_",
			"pack.class.func.owningAddressor : (a: Swift.Int, b: Swift.Int) -> (c: Swift.Int)",
		},
		{
			"_TFCC4pack6classA6classBlO4FuncFFS0_SSS1_",
			"pack.classA.classB.Func.owningAddressor : ((pack.classA) -> Swift.String) -> pack.classA.classB",
		},
		{
			"_TFC1Q1Q1QFTFSiSbSf_TSi1xFSbSf_",
			"Q.Q.Q((Swift.Int) -> Swift.Bool, Swift.Float) -> (Swift.Int, x: (Swift.Bool) -> Swift.Float)",
		},
		{
			"_TFC4Pack5class4FuncFTSiSbSS_TSSSiSb_",
			"Pack.class.Func(Swift.Int, Swift.Bool, Swift.String) -> (Swift.String, Swift.Int, Swift.Bool)",
		},
		{
			"_TFC4test5class4funcfT1QSS1bSi_TfSbSi_",
			"test.class.func(Q: Swift.String, b: Swift.Int) -> ((Swift.Bool) -> Swift.Int)",
		},
		{
			"_TFCC4test6classA6classB4funcfFFT1QSS1bSi_TfSbSi_TBtBO_TSQSRSV_",
			"test.classA.classB.func(((Q: Swift.String, b: Swift.Int) -> ((Swift.Bool) -> Swift.Int)) -> (Builtin.SILToken, Builtin.UnknownObject)) -> (Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeBufferPointer, Swift.UnsafeRawPointer)",
		},
		{
			"_TFCCC4test6classA6classB6classC4funcfFFFT1QSS1bSi_TfSbSiS2__TBbBo_TSQSpSu_TBw_",
			"test.classA.classB.classC.func((((Q: Swift.String, b: Swift.Int) -> ((Swift.Bool) -> Swift.Int, test.classA.classB.classC)) -> (Builtin.BridgeObject, Builtin.NativeObject)) -> (Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeMutablePointer, Swift.UInt)) -> (Builtin.Word)",
		},
		{
			"_TFCCCCC4Pack6classA6classB6classC6classD6classEG4FuncfffT1kSS2RxSb2AQSf_SST1QSS_TSSBBSQSR_",
			"Pack.classA.classB.classC.classD.classE.Func.getter : (((k: Swift.String, Rx: Swift.Bool, AQ: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.String, Builtin.UnsafeValueBuffer, Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeBufferPointer)",
		},
		{
			"_TFCCCC4Pack6classA6classB6classC6classDlp4FuncfffFT1kBw2RxSb2AWSf_SST1QSS_TSQ1wfSVSi_TSb1dfSiSf_",
			"Pack.classA.classB.classC.classD.Func.nativePinningAddressor : ((((k: Builtin.Word, Rx: Swift.Bool, AW: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.ImplicitlyUnwrappedOptional, w: (Swift.UnsafeRawPointer) -> Swift.Int)) -> (Swift.Bool, d: (Swift.Int) -> Swift.Float)",
		},
		{
			"_TFCC4pack6classA6classB4funcfCCC1a1b1c1dFSiSi",
			"pack.classA.classB.func(a.b.c.d) -> (Swift.Int) -> Swift.Int",
		},
		{
			"_TFCCCC4test6classA6classB6classC6classDao4funcfFFFFT1QSS1bSi_TfSbSiS2__TBbBo_TSQSpSR_TSC_TBB_",
			"test.classA.classB.classC.classD.func.nativeOwningMutableAddressor : (((((Q: Swift.String, b: Swift.Int) -> ((Swift.Bool) -> Swift.Int, test.classA.classB.classC)) -> (Builtin.BridgeObject, Builtin.NativeObject)) -> (Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeMutablePointer, Swift.UnsafeBufferPointer)) -> (__C_Synthesized)) -> (Builtin.UnsafeValueBuffer)",
		},
		{
			"_TFCCCC4test6classA6classB6classC6classDau4funcfFFFT1GSS1HSi_TfSbSiS3__TBbBo_T1QFSQSpSR_T1SSC_",
			"test.classA.classB.classC.classD.func.unsafeMutableAddressor : ((((G: Swift.String, H: Swift.Int) -> ((Swift.Bool) -> Swift.Int, test.classA.classB.classC.classD)) -> (Builtin.BridgeObject, Builtin.NativeObject)) -> (Q: (Swift.ImplicitlyUnwrappedOptional) -> Swift.UnsafeMutablePointer, Swift.UnsafeBufferPointer)) -> (S: __C_Synthesized)",
		},
		{
			"_TFCCCCC4Pack5iFirD5txcSk7QtfkWVy6jjaBfP6civmRIao4FuncfffT1kSS2RxSb2AQSf_SST1QSS_TSuSq_",
			"Pack.iFirD.txcSk.QtfkWVy.jjaBfP.civmRI.Func.nativeOwningMutableAddressor : (((k: Swift.String, Rx: Swift.Bool, AQ: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.UInt, Swift.Optional)",
		},
		{
			"_TFCCCCC4Pack6classA6classB6classC6classD6classIlp4FuncfffT1kBB2RxSb2AQSf_SST1QSS_TSQSV_",
			"Pack.classA.classB.classC.classD.classI.Func.nativePinningAddressor : (((k: Builtin.UnsafeValueBuffer, Rx: Swift.Bool, AQ: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeRawPointer)",
		},
		{
			"_TFCCCCC4Pack6classA6classB6classC6classD6classIlp4FuncfffFT1kBw2RxSb2AWSf_SST1QSS_TSQ1wfSVSi_Tt_1vBpBo_",
			"Pack.classA.classB.classC.classD.classI.Func.nativePinningAddressor : ((((k: Builtin.Word, Rx: Swift.Bool, AW: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.ImplicitlyUnwrappedOptional, w: (Swift.UnsafeRawPointer) -> Swift.Int)) -> ((), v: Builtin.RawPointer, Builtin.NativeObject)",
		},
		{
			"_TFCCCCC4Pack6classA6classB6classC6classD6civmRIao4FuncffFfT1kSS2RxSb2AQSf_SST1QSS_TSSBBSQSR_TSiSQ_",
			"Pack.classA.classB.classC.classD.civmRI.Func.nativeOwningMutableAddressor : ((((k: Swift.String, Rx: Swift.Bool, AQ: Swift.Float) -> Swift.String) -> (Q: Swift.String)) -> (Swift.String, Builtin.UnsafeValueBuffer, Swift.ImplicitlyUnwrappedOptional, Swift.UnsafeBufferPointer)) -> (Swift.Int, Swift.ImplicitlyUnwrappedOptional)",
		},
		{
			"_TFCCCOC4Pack9GbbKnBvMy6DNENum8XMogTAEc6VNXOSj9SAtlayOxgW4FuncFFSbT1WS4_S_1CSb3EjpS0_1eSP3wrcSV_TBB_",
			"Pack.GbbKnBvMy.DNENum.XMogTAEc.VNXOSj.SAtlayOxg.Func.didset : ((Swift.Bool) -> (W: Pack.GbbKnBvMy.DNENum.XMogTAEc.VNXOSj.SAtlayOxg, Pack, C: Swift.Bool, Ejp: Pack.GbbKnBvMy, e: Swift.UnsafePointer, wrc: Swift.UnsafeRawPointer)) -> (Builtin.UnsafeValueBuffer)",
		},
		{
			"_TFCCCOC4Pack9GbbKnBvMy6DNENum8XMogTAEc6VNXOSj9SAtlayOxgW4FuncFSbT1WS4_S_1CSb3EjpS0_1eS0_3wrcS1__",
			"Pack.GbbKnBvMy.DNENum.XMogTAEc.VNXOSj.SAtlayOxg.Func.didset : (Swift.Bool) -> (W: Pack.GbbKnBvMy.DNENum.XMogTAEc.VNXOSj.SAtlayOxg, Pack, C: Swift.Bool, Ejp: Pack.GbbKnBvMy, e: Pack.GbbKnBvMy, wrc: Pack.GbbKnBvMy.DNENum)",
		},
		{
			"_TFVCVVV4Pack8uLrpjlQH6dQDGbn5AZxTP8TJYmYUQp9lFoswzaads4FuncfSbt1YSb3cltS4_3ElYSi2MlS_2KUSb1pSb_",
			"Pack.uLrpjlQH.dQDGbn.AZxTP.TJYmYUQp.lFoswzaad.Func.setter : (Swift.Bool) -> (Y: Swift.Bool, clt: Pack.uLrpjlQH.dQDGbn.AZxTP.TJYmYUQp.lFoswzaad, ElY: Swift.Int, Ml: Pack, KU: Swift.Bool, p: Swift.Bool...)",
		},
		{
			"_TFVOCOV4Pack9IoLMpJzFP7qFjftmK6wWZqgE5MvxoR6SFZDdwlp4Funcft1IS3__tSr1jSSBt_",
			"Pack.IoLMpJzFP.qFjftmK.wWZqgE.MvxoR.SFZDdw.Func.nativePinningAddressor : (I: Pack.IoLMpJzFP.qFjftmK.wWZqgE.MvxoR...) -> (Swift.UnsafeMutableBufferPointer, j: Swift.String, Builtin.SILToken...)",
		},
		{
			"_TFCCCCV4Pack5vHYDl7oHJPprZ7TlOZURF5zRWbG8TakstNMClo4FuncftSi2PZS3_1ZSo_T3InqS4_2IGS3_2meS3_S3_2TJSP_",
			"Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG.TakstNMC.Func.nativeOwningAddressor : (Swift.Int, PZ: Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG, Z: __C...) -> (Inq: Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG.TakstNMC, IG: Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG, me: Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG, Pack.vHYDl.oHJPprZ.TlOZURF.zRWbG, TJ: Swift.UnsafePointer)",
		},
		{
			"_TFCVCOV4Pack7VRJxiBW9rVugCcwFf7tOdwiJn7cWhBiEx5XlgGSlO4FuncfFTSSSS_TSS_TSSSiSb_",
			"Pack.VRJxiBW.rVugCcwFf.tOdwiJn.cWhBiEx.XlgGS.Func.owningAddressor : ((Swift.String, Swift.String) -> (Swift.String)) -> (Swift.String, Swift.Int, Swift.Bool)",
		},
		{
			"_TFVVCVC4Pack7xCltHAS8UHRnscrf8EGmGEIDl8UMcRoACV5teKLvaO4FuncFT_t_",
			"Pack.xCltHAS.UHRnscrf.EGmGEIDl.UMcRoACV.teKLv.Func.owningMutableAddressor : () -> ()",
		},
		{
			"_TFOVCCO4Pack8vrahdHBh5QTgek9llzizyXPl9yDLktLAxh8FUmJiZablp4FuncfftBo_t_BB",
			"Pack.vrahdHBh.QTgek.llzizyXPl.yDLktLAxh.FUmJiZab.Func.nativePinningAddressor : ((Builtin.NativeObject...) -> ()) -> Builtin.UnsafeValueBuffer",
		},
		{
			"_TFCC4pack6classa6classbao4funcFFT1aSi1iBB_TSuSSSS_TBB_",
			"pack.classa.classb.func.nativeOwningMutableAddressor : ((a: Swift.Int, i: Builtin.UnsafeValueBuffer) -> (Swift.UInt, Swift.String, Swift.String)) -> (Builtin.UnsafeValueBuffer)",
		},
		{
			"_TFVCOOO4Pack8ULyCuctv9BKjaoArVO9xRfwpufje6XAmRrR5JEpebau4FuncFTS1__TSSS0_BB_",
			"Pack.ULyCuctv.BKjaoArVO.xRfwpufje.XAmRrR.JEpeb.Func.unsafeMutableAddressor : (Pack.ULyCuctv.BKjaoArVO) -> (Swift.String, Pack.ULyCuctv, Builtin.UnsafeValueBuffer)",
		},
		{
			"_TFOVOOC4Pack8yPHoAJqn7ogGhcvH6lYcjRJ6WGldTo8OcVTqVWBw4FuncfTSbSbSS2tcS3_S3_3aGQS1__TSiS_3rDYS3_1hS__",
			"Pack.yPHoAJqn.ogGhcvH.lYcjRJ.WGldTo.OcVTqVWB.Func.willset : (Swift.Bool, Swift.Bool, Swift.String, tc: Pack.yPHoAJqn.ogGhcvH.lYcjRJ.WGldTo, Pack.yPHoAJqn.ogGhcvH.lYcjRJ.WGldTo, aGQ: Pack.yPHoAJqn.ogGhcvH) -> (Swift.Int, Pack, rDY: Pack.yPHoAJqn.ogGhcvH.lYcjRJ.WGldTo, h: Pack)",
		},
		{
			"_TFVOOCV4Pack7HhGMiSf5DjdjA8zlMZmXbz8VvAZxuRL6EkQHCig4FuncftS1_2eoSf2ZkS_2HGSb1MSb_TS0_2bxSf2cwSS2txBo_",
			"Pack.HhGMiSf.DjdjA.zlMZmXbz.VvAZxuRL.EkQHCi.Func.getter : (Pack.HhGMiSf.DjdjA, eo: Swift.Float, Zk: Pack, HG: Swift.Bool, M: Swift.Bool...) -> (Pack.HhGMiSf, bx: Swift.Float, cw: Swift.String, tx: Builtin.NativeObject)",
		},
		{
			"_TFCOCVV4Pack7ORBSacp9iKxTJZmEX6KEDgAV6lRQerP9HfhXoMrSyw4FuncfS2_T1XSf2GsS3_3dPJS4_3iQJSb_",
			"Pack.ORBSacp.iKxTJZmEX.KEDgAV.lRQerP.HfhXoMrSy.Func.willset : (Pack.ORBSacp.iKxTJZmEX.KEDgAV) -> (X: Swift.Float, Gs: Pack.ORBSacp.iKxTJZmEX.KEDgAV.lRQerP, dPJ: Pack.ORBSacp.iKxTJZmEX.KEDgAV.lRQerP.HfhXoMrSy, iQJ: Swift.Bool)",
		},
		{
			"_TFOOVCC4Pack5xysGw8IjlvQmzd7kDzDMjQ5OczyY6XkDDtMw4Funcft3JPZS3_1VS4_1nSS1lSS1qSS_T2BVSS3ieuSbSS1OS2_1hS1_S1__",
			"Pack.xysGw.IjlvQmzd.kDzDMjQ.OczyY.XkDDtM.Func.willset : (JPZ: Pack.xysGw.IjlvQmzd.kDzDMjQ.OczyY, V: Pack.xysGw.IjlvQmzd.kDzDMjQ.OczyY.XkDDtM, n: Swift.String, l: Swift.String, q: Swift.String...) -> (BV: Swift.String, ieu: Swift.Bool, Swift.String, O: Pack.xysGw.IjlvQmzd.kDzDMjQ, h: Pack.xysGw.IjlvQmzd, Pack.xysGw.IjlvQmzd)",
		},
		{
			"_TFOOCOV4Pack9NLMaFIUHe5yeWjw9duVKJWELT6hhWTFd5hjhbCG4FuncfT2umSb3BzUS_2ByS3_Sf2OpSb_T2uRS_3rRcS4_3pmyS1_Sb3cXqSQ_",
			"Pack.NLMaFIUHe.yeWjw.duVKJWELT.hhWTFd.hjhbC.Func.getter : (um: Swift.Bool, BzU: Pack, By: Pack.NLMaFIUHe.yeWjw.duVKJWELT.hhWTFd, Swift.Float, Op: Swift.Bool) -> (uR: Pack, rRc: Pack.NLMaFIUHe.yeWjw.duVKJWELT.hhWTFd.hjhbC, pmy: Pack.NLMaFIUHe.yeWjw, Swift.Bool, cXq: Swift.ImplicitlyUnwrappedOptional)",
		},
		{
			"_TFVCCOC4Pack8YTppvaIs7QddEyPH6BoHgqM6eVKUHI8VhITsOZKao4FuncfTS2_2TvS3_2ZRS2_S1_1qS_SS_Bw",
			"Pack.YTppvaIs.QddEyPH.BoHgqM.eVKUHI.VhITsOZK.Func.nativeOwningMutableAddressor : (Pack.YTppvaIs.QddEyPH.BoHgqM, Tv: Pack.YTppvaIs.QddEyPH.BoHgqM.eVKUHI, ZR: Pack.YTppvaIs.QddEyPH.BoHgqM, Pack.YTppvaIs.QddEyPH, q: Pack, Swift.String) -> Builtin.Word",
		},
		{
			"_TFCVOVO4Pack6DMlRIQ5qpApc5eBYLG9GbbuAVJLb6FXRNgQg4FuncfT1qS2_3IqzS4__T3mNES3_1qS0_2HkS1_3sZQSb2UGSS1RS__",
			"Pack.DMlRIQ.qpApc.eBYLG.GbbuAVJLb.FXRNgQ.Func.getter : (q: Pack.DMlRIQ.qpApc.eBYLG, Iqz: Pack.DMlRIQ.qpApc.eBYLG.GbbuAVJLb.FXRNgQ) -> (mNE: Pack.DMlRIQ.qpApc.eBYLG.GbbuAVJLb, q: Pack.DMlRIQ, Hk: Pack.DMlRIQ.qpApc, sZQ: Swift.Bool, UG: Swift.String, R: Pack)",
		},
		{
			"_TFC4pack5class4funcFTFSbSfSf_TSi1xFSbSf_",
			"pack.class.func((Swift.Bool) -> Swift.Float, Swift.Float) -> (Swift.Int, x: (Swift.Bool) -> Swift.Float)",
		},
		{
			"_TFCCCCC4Pack9LBmkxqtje8HkdiTBuG7rSQfceV7vLMMUYg6BUxKUWm4FuncftSf3zolSf2GbS4_2SES0__tS2_3gdfS3_3JBCSi_",
			"Pack.LBmkxqtje.HkdiTBuG.rSQfceV.vLMMUYg.BUxKUW.Func.materializeForSet : (Swift.Float, zol: Swift.Float, Gb: Pack.LBmkxqtje.HkdiTBuG.rSQfceV.vLMMUYg.BUxKUW, SE: Pack.LBmkxqtje...) -> (Pack.LBmkxqtje.HkdiTBuG.rSQfceV, gdf: Pack.LBmkxqtje.HkdiTBuG.rSQfceV.vLMMUYg, JBC: Swift.Int...)",
		},
		{
			"_TFC4pack5class4funcFT1qSSSf_TSb1pfSiSiSb_",
			"pack.class.func(q: Swift.String, Swift.Float) -> (Swift.Bool, p: (Swift.Int) -> Swift.Int, Swift.Bool)",
		},
		{
			"_TFC4test5class4funcFT1aSS1bSb1cSi_FTSiSi_Si",
			"test.class.func(a: Swift.String, b: Swift.Bool, c: Swift.Int) -> (Swift.Int, Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4test5class4funcFT1aSf1bSi1cSi_FT1dSS1eSi_Si",
			"test.class.func(a: Swift.Float, b: Swift.Int, c: Swift.Int) -> (d: Swift.String, e: Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4test5class4funcFT1aSi1bSb_SS",
			"test.class.func(a: Swift.Int, b: Swift.Bool) -> Swift.String",
		},
		{
			"_TFVCCCC4Pack6classA6classB6classC6classD6classFap4FuncfT1KS0_1CSi2fmS1_2MxSSSb_tSb1MSf3xdoSS1GS1_3lKiSC_",
			"Pack.classA.classB.classC.classD.classF.Func.nativePinningMutableAddressor : (K: Pack.classA, C: Swift.Int, fm: Pack.classA.classB, Mx: Swift.String, Swift.Bool) -> (Swift.Bool, M: Swift.Float, xdo: Swift.String, G: Pack.classA.classB, lKi: __C_Synthesized...)",
		},
		{
			"_TFVCVOO4Pack6sCUjVy7ScIFZau9XawJzRIqf9swlnLSlJp9VAjeugXpkau4FuncFT1bS1_1US2_3rfsS1_1YS1_1US1_1dSS_S_",
			"Pack.sCUjVy.ScIFZau.XawJzRIqf.swlnLSlJp.VAjeugXpk.Func.unsafeMutableAddressor : (b: Pack.sCUjVy.ScIFZau, U: Pack.sCUjVy.ScIFZau.XawJzRIqf, rfs: Pack.sCUjVy.ScIFZau, Y: Pack.sCUjVy.ScIFZau, U: Pack.sCUjVy.ScIFZau, d: Swift.String) -> Pack",
		},
		{
			"_TFC1q1q1qfTSiSbSf_TSifSiSi_",
			"q.q.q(Swift.Int, Swift.Bool, Swift.Float) -> (Swift.Int, (Swift.Int) -> Swift.Int)",
		},
		{
			"_TFC4test5class4funcFT1aSS1bSb1cSi_FSiSi",
			"test.class.func(a: Swift.String, b: Swift.Bool, c: Swift.Int) -> (Swift.Int) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFTSiSb_SSSSSi",
			"pack.class.Func(((Swift.Int, Swift.Bool) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFOCCVO4Pack6classA6classB6classC6classD6classFap4FuncfffTSiSp_TfSCSi_TSiSSS2__t1aSiS0_S1__",
			"Pack.classA.classB.classC.classD.classF.Func.nativePinningMutableAddressor : (((Swift.Int, Swift.UnsafeMutablePointer) -> ((__C_Synthesized) -> Swift.Int)) -> (Swift.Int, Swift.String, Pack.classA.classB.classC)) -> (a: Swift.Int, Pack.classA, Pack.classA.classB...)",
		},
		{
			"_TFC4pack5class4FuncFFFSbTSSSb_SSSi",
			"pack.class.Func(((Swift.Bool) -> (Swift.String, Swift.Bool)) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFTSiSb_TSSSb_SSSi",
			"pack.class.Func(((Swift.Int, Swift.Bool) -> (Swift.String, Swift.Bool)) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFTSiFSiSb_SSSSSi",
			"pack.class.Func(((Swift.Int, (Swift.Int) -> Swift.Bool) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFTSiFTSiSS_Sb_SSSSSi",
			"pack.class.Func(((Swift.Int, (Swift.Int, Swift.String) -> Swift.Bool) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFTSiFTSiFSiSb_Sb_SSSSSi",
			"pack.class.Func(((Swift.Int, (Swift.Int, (Swift.Int) -> Swift.Bool) -> Swift.Bool) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFT1dSiFT1aSiFSiSb_Sb_SSSSSi",
			"pack.class.Func(((d: Swift.Int, (a: Swift.Int, (Swift.Int) -> Swift.Bool) -> Swift.Bool) -> Swift.String) -> Swift.String) -> Swift.Int",
		},
		{
			"_TFC4pack5class4FuncFFFT1aSiFT1bSi1cSS_Si_SbSSSi",
			"pack.class.Func(((a: Swift.Int, (b: Swift.Int, c: Swift.String) -> Swift.Int) -> Swift.Bool) -> Swift.String) -> Swift.Int",
		},
	}
	for _, test := range tests {
		got, _ := ToString(test.input)
		if got != test.want {
			t.Logf("[fail] want %s ,get %s\n", test.want, got)
		}
	}
}
