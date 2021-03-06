
package main 

import (
	"fmt"
	"log"
	"errors"
	"unsafe"
	"strings"
	"syscall"
	"io/ioutil"
	"encoding/json"
)


// json data from https://github.com/telllpu/WhatAV
var av_json_data string = `{
	"ALYac":{
		"processes":["aylaunch.exe","ayupdate2.exe","AYRTSrv.exe","AYAgent.exe"],
		"url":"https://en.estsecurity.com/"},
	"AVG":{
		"processes":["AVGSvc.exe","AVGUI.exe","avgwdsvc.exe","avg.exe","avgaurd.exe","avgemc.exe","avgrsx.exe","avgserv.exe","avgw.exe"],
		"url":"https://www.avg.com/"},
	"Acronis":{
		"processes":["arsm.exe","acronis_license_service.exe"],
		"url":"https://www.acronis.com/"},
	"Ad-Aware":{
		"processes":["AdAwareService.exe","Ad-Aware.exe","AdAware.exe"],
		"url":"https://www.adaware.com/"},
	"AhnLab-V3":{
		"processes":["patray.exe","V3Svc.exe"],
		"url":"https://global.ahnlab.com/site/main.do"},
	"Arcabit":{
		"processes":["arcavir.exe","arcadc.exe","ArcaVirMaster.exe","ArcaMainSV.exe","ArcaTasksService.exe"],
		"url":"https://www.arcabit.pl"},
	"Avast":{
		"processes":["ashDisp.exe","AvastUI.exe","AvastSvc.exe","AvastBrowser.exe","AfwServ.exe"],
		"url":"https://www.avast.com"},
	"Avira AntiVirus":{
		"processes":["avcenter.exe","avguard.exe","avgnt.exe","sched.exe"],
		"url":"https://www.avira.com/"},
	"百度杀毒":{
		"processes":["BaiduSdSvc.exe","BaiduSdTray.exe","BaiduSd.exe","bddownloader.exe","baiduansvx.exe"],
		"url":"https://anquan.baidu.com/"},
	"BitDefender":{
		"processes":["Bdagent.exe","BitDefenderCom.exe","vsserv.exe","bdredline.exe","bdservicehost.exe"],
		"url":"http://www.bitdefender.com/"},
	"Bkav":{
		"processes":["BKavService.exe","Bka.exe","BkavUtil.exe","BLuPro.exe"],
		"url":"https://www.bkav.com/"},
	"CAT-QuickHeal":{
		"processes":["QUHLPSVC.exe","onlinent.exe","sapissvc.exe","scanwscs.exe"],
		"url":"https://www.quickheal.com/"},
	"CMC":{
		"processes":["CMCTrayIcon.exe"],
		"url":"https://cmccybersecurity.com/"},
	"ClamAV":{
		"processes":["freshclam.exe"],
		"url":"https://www.clamav.net"},
	"Comodo":{
		"processes":["cpf.exe","cavwp.exe","ccavsrv.exe","cmdvirth.exe"],
		"url":"https://www.comodo.com"},
	"CrowdStrike Falcon":{
		"processes":["csfalconservice.exe","CSFalconContainer.exe"],
		"url":"https://www.crowdstrike.com"},
	"Cybereason":{
		"processes":["CybereasonRansomFree.exe","CybereasonRansomFreeServiceHost.exe","CybereasonAV.exe"],
		"url":"https://www.cybereason.com/"},
	"Cylance":{
		"processes":["CylanceSvc.exe"],
		"url":"https://www.cylance.com"},
	"Cyren":{
		"processes":["vsedsps.exe","vseamps.exe","vseqrts.exe"],
		"url":"http://www.cyren.com/"},
	"DrWeb":{
		"processes":["drwebcom.exe","spidernt.exe","drwebscd.exe","drweb32w.exe","dwengine.exes"],
		"url":"https://www.drweb.com/"},
	"ESET-NOD32":{
		"processes":["egui.exe","ecls.exe","ekrn.exe","eguiProxy.exe"],
		"url":"https://www.eset.com/us/home/antivirus/"},
	"Emsisoft":{
		"processes":["a2cmd.exe","a2guard.exe"],
		"url":"https://www.emsisoft.com/"},
	"Endgame":{
		"processes":["endgame.exe"],
		"url":"https://www.endgame.com/"},
	"F-Prot":{
		"processes":["F-PROT.exe","FProtTray.exe","FPAVServer.exe","f-stopw.exe","f-prot95.exe","f-agnt95.exe"],
		"url":"http://f-prot.com/"},
	"F-Secure":{
		"processes":["f-secure.exe","fssm32.exe","Fsorsp64.exe","fsavgui.exe","fameh32.exe","fch32.exe","fih32.exe","fnrb32.exe","fsav32.exe","fsma32.exe","fsmb32.exe"],
		"url":"https://www.f-secure.com"},
	"FireEye":{
		"processes":["xagtnotif.exe","xagt.exe"],
		"url":"https://www.fireeye.com/"},
	"Fortinet":{
		"processes":["FortiClient.exe","FortiTray.exe","FortiScand.exe"],
		"url":"https://fortiguard.com/"},
	"GData":{
		"processes":["AVK.exe","avkcl.exe","avkpop.exe","avkservice.exe"],
		"url":"https://www.gdatasoftware.com/"},
	"Ikarus":{
		"processes":["guardxservice.exe","guardxkickoff.exe"],
		"url":"https://www.ikarussecurity.com/"},
	"江民":{
		"processes":["KVFW.exe","KVsrvXP.exe","KVMonXP.exe","KVwsc.exe"],
		"url":"https://www.jiangmin.com/"},
	"K7AntiVirus":{
		"processes":["K7TSecurity.exe","K7TSMain.Exe","K7TSUpdT.exe"],
		"url":"http://viruslab.k7computing.com/"},
	"Kaspersky":{
		"processes":["avp.exe","avpcc.exe","avpm.exe","kavpf.exe"],
		"url":"https://www.kaspersky.com"},
	"Kingsoft":{
		"processes":["kxetray.exe","ksafe.exe","KSWebShield.exe","kpfwtray.exe","KWatch.exe","KSafeSvc.exe","KSafeTray.exe"],
		"url":"http://www.duba.net/"},
	"Max Secure Software":{
		"processes":["SDSystemTray.exe","MaxRCSystemTray.exe","RCSystemTray.exe"],
		"url":"https://www.maxpcsecure.com/"},
	"Malwarebytes":{
		"processes":["MalwarebytesPortable.exe","Mbae.exe","MBAMIService.exe","mbamdor.exe"],
		"url":"http://www.malwarebytes.org/"},
	"McAfee":{
		"processes":["Mcshield.exe","Tbmon.exe","Frameworkservice.exe","firesvc.exe","firetray.exe","hipsvc.exe","mfevtps.exe","mcafeefire.exe","shstat.exe","vstskmgr.exe","engineserver.exe","alogserv.exe","avconsol.exe","cmgrdian.exe","cpd.exe","mcmnhdlr.exe","mcvsshld.exe","mcvsrte.exe","mghtml.exe","mpfservice.exe","mpfagent.exe","mpftray.exe","vshwin32.exe","vsstat.exe","guarddog.exe"],
		"url":"https://www.mcafee.com/en-us"},
	"Microsoft Security Essentials":{
		"processes":["MsMpEng.exe","mssecess.exe","emet_service.exe","drwatson.exe"],
		"url":"https://support.microsoft.com/en-us/help/17150/windows-7-what-is-microsoft-security-essentials"},
	"NANO-Antivirus":{
		"processes":["nanoav.exe","nanoav64.exe","nanoreport.exe","nanoreportc.exe","nanoreportc64.exe","nanorst.exe","nanosvc.exe"],
		"url":"https://nano-av.com/"},
	"a-squared free":{
		"processes":["a2guard.exe","a2free.exe","a2service.exe"],
		"url":"https://baike.baidu.com/item/a-squared%20Free/481873?fr=aladdin"},
	"Palo Alto Networks":{
		"processes":["PanInstaller.exe"],
		"url":"https://www.paloaltonetworks.com/"},
	"Panda Security":{
		"processes":["remupd.exe","apvxdwin.exe","pavproxy.exe","pavsched.exe"],
		"url":"https://www.pandasecurity.com/"},
	"Qihoo-360":{
		"processes":["360sd.exe","360tray.exe","ZhuDongFangYu.exe","360rp.exe","360safe.exe","360safebox.exe","QHActiveDefense.exe"],
		"url":"https://sd.360.cn/"},
	"Rising":{
		"processes":["RavMonD.exe","rfwmain.exe","RsMgrSvc.exe"],
		"url":"http://antivirus.rising.com.cn/"},
	"SUPERAntiSpyware":{
		"processes":["superantispyware.exe","sascore.exe","SAdBlock.exe","sabsvc.exe"],
		"url":"http://www.superadblocker.com/"},
	"SecureAge APEX":{
		"processes":["UniversalAVService.exe","EverythingServer.exe","clamd.exe"],
		"url":"https://www.secureage.com/"},
	"SentinelOne (Static ML)":{
		"processes":[],
		"url":"https://www.sentinelone.com/"},
	"Sophos AV":{
		"processes":["SavProgress.exe","SophosUI.exe","SophosFS.exe","SophosHealth.exe","SophosSafestore64.exe","SophosCleanM.exe","icmon.exe"],
		"url":"https://www.sophos.com/"},
	"Symantec":{
		"processes":["ccSetMgr.exe","ccapp.exe","vptray.exe","ccpxysvc.exe","cfgwiz.exe","smc.exe","symproxysvc.exe","vpc32.exe","lsetup.exe","luall.exe","lucomserver.exe","sbserv.exe"],
		"url":"http://www.symantec.com/"},
	"TACHYON":{
		"processes":[],
		"url":"https://www.tachyonlab.com/en/index.html"},
	"Tencent":{
		"processes":["QQPCRTP.exe","QQPCTray.exe","QQPCMgr.exe"],
		"url":"https://guanjia.qq.com"},
	"TotalDefense":{
		"processes":["AMRT.exe","SWatcherSrv.exe","Prd.ManagementConsole.exe"],
		"url":"https://www.totaldefense.com"},
	"Trapmine":{
		"processes":["TrapmineEnterpriseService.exe","TrapmineEnterpriseConfig.exe","TrapmineDeployer.exe","TrapmineUpgradeService.exe"],
		"url":"https://trapmine.com/"},
	"TrendMicro":{
		"processes":["TMBMSRV.exe","ntrtscan.exe","Pop3Trap.exe","WebTrap.exe"],
		"url":"http://careers.trendmicro.com.cn/"},
	"VIPRE":{
		"processes":["SBAMSvc.exe","VipreEdgeProtection.exe","SBAMTray.exe"],
		"url":"https://www.vipre.com"},
	"ViRobot":{
		"processes":["vrmonnt.exe","vrmonsvc.exe","Vrproxyd.exe"],
		"url":"http://www.hauri.net/"},
	"Webroot":{
		"processes":["npwebroot.exe","WRSA.exe","spysweeperui.exe"],
		"url":"https://www.webroot.com/us/en"},
	"Yandex":{
		"processes":["Yandex.exe","YandexDisk.exe","yandesk.exe"],
		"url":"https://yandex.com/support/common/security/antiviruses-free.html"},
	"Zillya":{
		"processes":["zillya.exe","ZAVAux.exe","ZAVCore.exe"],
		"url":"https://zillya.com"},
	"ZoneAlarm":{
		"processes":["vsmon.exe","zapro.exe","zonealarm.exe"],
		"url":"https://www.zonealarm.com/"},
	"Zoner":{
		"processes":["ZPSTray.exe"],
		"url":"https://zonerantivirus.com/"},
	"eGambit":{
		"processes":["dasc.exe","dastray.exe","memscan64.exe","dastray.exe"],
		"url":"https://egambit.app/en/"},
	"eScan":{
		"processes":["consctl.exe","mwaser.exe","avpmapp.exe"],
		"url":"https://www.escanav.com/"},
	"Lavasoft":{
		"processes":["AAWTray.exe","LavasoftTcpService.exe","AdAwareTray.exe","WebCompanion.exe","WebCompanionInstaller.exe","adawarebp.exe"],
		"url":"https://www.lavasoft.com/"},
	"The Cleaner":{
		"processes":["cleaner8.exe"],
		"url":""},
	"VBA32":{
		"processes":["vba32lder.exe"],
		"url":"http://www.anti-virus.by/en/index.shtml"},
	"Mongoosa":{
		"processes":["MongoosaGUI.exe","mongoose.exe"],
		"url":"https://www.securitymongoose.com/"},
	"Coranti2012":{
		"processes":["CorantiControlCenter32.exe"],
		"url":"https://www.coranti.com"},
	"UnThreat":{
		"processes":["UnThreat.exe","utsvc.exe"],
		"url":"https://softplanet.com/UnThreat-AntiVirus"},
	"Shield Antivirus":{
		"processes":["CKSoftShiedAntivirus4.exe","shieldtray.exe"],
		"url":"https://shieldapps.com/supportmain/shield-antivirus-support/"},
	"VIRUSfighter":{
		"processes":["AVWatchService.exe","vfproTray.exe"],
		"url":"https://www.spamfighter.com/VIRUSfighter/"},	
	"Immunet":{
		"processes":["iptray.exe"],
		"url":"https://www.immunet.com/index"},
	"PSafe":{
		"processes":["PSafeSysTray.exe","PSafeCategoryFinder.exe","psafesvc.exe"],
		"url":"https://www.psafe.com/"},
	"nProtect":{
		"processes":["nspupsvc.exe","Npkcmsvc.exe","npnj5Agent.exe"],
		"url":"http://nos.nprotect.com/"},
	"Spyware Terminator":{
		"processes":["SpywareTerminatorShield.exe","SpywareTerminator.exe"],
		"url":"http://www.spywareterminator.com/Default.aspx"},
	"Norton":{
		"processes":["ccSvcHst.exe","rtvscan.exe","ccapp.exe","NPFMntor.exe","ccRegVfy.exe","vptray.exe","iamapp.exe","nav.exe","navapw32.exe","navapsvc.exe","nisum.exe","nmain.exe","nprotect.exe"],
		"url":"https://us.norton.com/"},
	"可牛杀毒":{
		"processes":["knsdtray.exe"],
		"url":"https://baike.baidu.com/item/%E5%8F%AF%E7%89%9B%E5%85%8D%E8%B4%B9%E6%9D%80%E6%AF%92%E8%BD%AF%E4%BB%B6"},
	"流量矿石":{
		"processes":["Miner.exe"],
		"url":"https://jiaoyi.yunfan.com/"},
	"safedog":{
		"processes":["safedog.exe","SafeDogGuardCenter.exe","safedogupdatecenter.exe","safedogguardcenter.exe","SafeDogSiteIIS.exe","SafeDogTray.exe","SafeDogServerUI.exe"],
		"url":"http://www.safedog.cn/"},
	"木马克星":{
		"processes":["parmor.exe","Iparmor.exe"],
		"url":"https://baike.baidu.com/item/%E6%9C%A8%E9%A9%AC%E5%85%8B%E6%98%9F/2979824?fr=aladdin"},
	"贝壳云安全":{
		"processes":["beikesan.exe"],
		"url":""},
	"木马猎手":{
		"processes":["TrojanHunter.exe"],
		"url":""},
	"巨盾网游安全盾":{
		"processes":["GG.exe"],
		"url":""},
	"绿鹰安全精灵":{
		"processes":["adam.exe"],
		"url":"https://baike.baidu.com/item/%E7%BB%BF%E9%B9%B0%E5%AE%89%E5%85%A8%E7%B2%BE%E7%81%B5"},
	"超级巡警":{
		"processes":["AST.exe"],
		"url":""},
	"墨者安全专家":{
		"processes":["ananwidget.exe"],
		"url":""},
	"风云防火墙":{
		"processes":["FYFireWall.exe"],
		"url":""},
	"微点主动防御":{
		"processes":["MPMon.exe"],
		"url":"http://www.micropoint.com.cn/"},
	"天网防火墙":{
		"processes":["pfw.exe"],
		"url":""},
	"D盾":{
		"processes":["D_Safe_Manage.exe","d_manage.exe"],
		"url":"http://www.d99net.net/"},
	"云锁":{
		"processes":["yunsuo_agent_service.exe","yunsuo_agent_daemon.exe"],
		"url":"https://www.yunsuo.com.cn/"},
	"护卫神":{
		"processes":["HwsPanel.exe","hws_ui.exe","hws.exe","hwsd.exe"],
		"url":"https://www.hws.com/"},
	"火绒":{
		"processes":["hipstray.exe","wsctrl.exe","usysdiag.exe"],
		"url":""},
	"网络病毒克星":{
		"processes":["WEBSCANX.exe"],
		"url":""},
	"SPHINX防火墙":{
		"processes":["SPHINX.exe"],
		"url":""},
	"Enhanced Mitigation Experience Toolkit":{
		"processes":["emet_agent.exe","emet_service.exe"],
		"url":"https://support.microsoft.com/en-us/help/2458544/the-enhanced-mitigation-experience-toolkit"},
	"H+BEDV Datentechnik GmbH":{
		"processes":["avwin.exe","avwupsrv.exe"],
		"url":"http://www.free-av.com/"},
	"IBM ISS Proventia":{
		"processes":["blackd.exe","rapapp.exe"],
		"url":""},
	"eEye Digital Security":{
		"processes":["eeyeevnt.exe","blink.exe"],
		"url":""},
	"TamoSoft":{
		"processes":["cv.exe","ent.exe"],
		"url":"https://www.tamos.com/"},
	"Kerio Personal Firewall":{
		"processes":["persfw.exe","wrctrl.exe"],
		"url":"http://www.kerio.com/"},
	"Simplysup":{
		"processes":["Trjscan.exe"],
		"url":"https://www.simplysup.com/"},
	"PC Tools AntiVirus":{
		"processes":["PCTAV.exe","pctsGui.exe"],
		"url":"http://www.pctools.com"},
	"VirusBuster Professional":{
		"processes":["vbcmserv.exe"],
		"url":"http://www.virusbuster.hu"},
	"ClamWin":{
		"processes":["ClamTray.exe","clamscan.exe"],
		"url":"http://www.clamwin.com/"},	
	"安天智甲":{
		"processes":["kxetray.exe","kscan.exe","AMediumManager.exe","kismain.exe"],
		"url":"https://antiy.cn/"}
}
`
const MAX_PATH = 260
type PROCESSENTRY32 struct {
	Size              uint32
	CntUsage          uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	CntThreads        uint32
	ParentProcessID   uint32
	PriorityClassBase int32
	Flags             uint32
	ExeFile           [MAX_PATH]uint16
}
func enumerate_process() ([]string, error) {
	fmt.Println("entering process enumerate...")
	
	modKernel32 := syscall.NewLazyDLL("kernel32.dll")
	procCreateToolhelp32Snapshot := modKernel32.NewProc("CreateToolhelp32Snapshot")
	procCloseHandle := modKernel32.NewProc("CloseHandle")
	procProcess32First := modKernel32.NewProc("Process32FirstW")
	procProcess32Next := modKernel32.NewProc("Process32NextW")
	handle, _, _ := procCreateToolhelp32Snapshot.Call(0x00000002, 0)
	if handle < 0 {
		fmt.Println(syscall.GetLastError())
		return  []string{}, syscall.GetLastError()
	}
	defer procCloseHandle.Call(handle)

	var entry PROCESSENTRY32
	entry.Size = uint32(unsafe.Sizeof(entry))
	ret, _, _ := procProcess32First.Call(handle, uintptr(unsafe.Pointer(&entry)))
	if ret == 0 {
		fmt.Errorf("Error retrieving process info.")
		return  []string{}, fmt.Errorf("Error retrieving process info.")
	}
	// fmt.Printf("first %d#######%v\n", ret, syscall.UTF16ToString(entry.ExeFile[:]))//这个默认是系统进程
		
	var results []string
	for {
		ret, _, _ := procProcess32Next.Call(handle, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}
		// fmt.Printf("进程ID：%d\t进程名：%s\n", entry.ProcessID, syscall.UTF16ToString(entry.ExeFile[:]))
		results = append(results, syscall.UTF16ToString(entry.ExeFile[:]))
	}

	return results, nil
}

type AV_info struct {
	Processes 	[]string  	`json:"processes"`
	Url 		string 		`json:"url"`
}

func check_av(all_process_list []string) (ret_str string, err error){
	fmt.Println("\nAnti-Virus process check...")
	count := 0

	var av map[string]AV_info
	err = json.Unmarshal([]byte(av_json_data), &av)
	if err != nil {
		log.Fatal("unmarshal err : ", err)
	}

	var results_str string 
	for key, kv := range av {
		for _, proc := range kv.Processes {
			proc = strings.ToLower(proc) //进程名变成小写
			for _, tmp_proc := range all_process_list {
				tmp_proc2 := strings.ToLower(tmp_proc)
				if tmp_proc2 == proc {
					fmt.Printf("%s Found : %s\n", key, tmp_proc)
					tmp_res := fmt.Sprintf("\t%s Found : %s\n", key, tmp_proc)
					results_str += tmp_res
					count += 1
				}
			}
		}
	}

	if count <= 0 {
		fmt.Println("\nNo AntiVirus Software Found !")
		return "", errors.New("毛都没找到")
	}

	return results_str, nil
}


//使用map去除重复元素
func trip_duplicate_map(original []string) (results []string) {
	tmp := map[string]byte{}
	for _, kv := range original {
		l := len(tmp)
		tmp[kv] = 0
		if len(tmp) != l {//加入map后，map长度变化，则元素不重复
			results = append(results, kv)
		}
	}

	return results
}

func main() {
	all_process_list, err := enumerate_process()
	if err != nil {
		log.Fatal("enum process failed:", err)
	}
	all_process_list = trip_duplicate_map(all_process_list)
	ret, err := check_av(all_process_list)
	if err != nil {
		ret = "检测结果：毛都没找到！"
	}else{
		ret = "检测结果：\n"+ret
	}

	ioutil.WriteFile("results.txt", []byte(ret), 0644)
}
