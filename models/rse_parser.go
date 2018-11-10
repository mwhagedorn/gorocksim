package models

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type RSEParser struct {
	Context EngineParser
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type EngineDatabase struct {
	XMLName xml.Name `xml:"engine-database"`
	Database EngineList `xml:"engine-list"`
}

type EngineList struct {
	XMLName xml.Name `xml:"engine-list"`
	Engines []RSEEngine `xml:"engine"`
}

type RSEEngine struct {
	Comments string `xml:"comments"`
	Code  string `xml:"code,attr"`
	Diameter float64 `xml:"dia,attr"`
	Length float64 `xml:"len,attr"`
	Delay string `xml:"delays,attr"`
	PropellantWeight float64 `xml:"propWt,attr"`
	EngineWeight float64 `xml:"initWt,attr"`
	Manufacturer string `xml:"mfg,:attr"`
	BurnTime float64 `xml:"burn-time,attr"`
	Data []RSEDataPoint `xml:"data>eng-data"`
}

type RSEDataPoint struct {
	TimeStamp float64 `xml:"t,attr"`
	Force float64 `xml:"f,attr"`
	Mass float64 `xml:"m,attr"`
}

// <engine-database>
// <engine-list>
// <engine FDiv="10" FFix="1" FStep="-1." Isp="71.7" Itot="2.32" Type="single-use" auto-calc-cg="1" auto-calc-mass="1" avgThrust="3.178" burn-time="0.73" cgDiv="10" cgFix="1" cgStep="-1." code="A8" delays="3,5" dia="18." exitDia="0." initWt="16.35" len="70." mDiv="10" mFix="1" mStep="-1." massFrac="20.18" mfg="Estes" peakThrust="9.73" propWt="3.3" tDiv="10" tFix="1" tStep="-1." throatDia="0.">
// <comments>Estes A8 RASP.ENG file made from NAR published data
// File produced October 3, 2000
// The total impulse, peak thrust, average thrust and burn time are
// the same as the averaged static test data on the NAR web site in
// the certification file. The curve drawn with these data points is as
// close to the certification curve as can be with such a limited
// number of points (32) allowed with wRASP up to v1.6.
// </comments>
// <data>
// <eng-data cg="35." f="0." m="3.3" t="0."/>
// <eng-data cg="35." f="0.512" m="3.28507" t="0.041"/>
// <eng-data cg="35." f="2.115" m="3.20474" t="0.084"/>
// <eng-data cg="35." f="4.358" m="3.0068" t="0.127"/>
// <eng-data cg="35." f="6.794" m="2.6975" t="0.166"/>
// <eng-data cg="35." f="8.588" m="2.41309" t="0.192"/>
// <eng-data cg="35." f="9.294" m="2.23506" t="0.206"/>
// <eng-data cg="35." f="9.73" m="1.96448" t="0.226"/>
// <eng-data cg="35." f="8.845" m="1.83238" t="0.236"/>
// <eng-data cg="35." f="7.179" m="1.70703" t="0.247"/>
// <eng-data cg="35." f="5.063" m="1.58515" t="0.261"/>
// <eng-data cg="35." f="3.717" m="1.48525" t="0.277"/>
// <eng-data cg="35." f="3.205" m="1.3425" t="0.306"/>
// <eng-data cg="35." f="2.884" m="1.14764" t="0.351"/>
// <eng-data cg="35." f="2.499" m="0.94092" t="0.405"/>
// <eng-data cg="35." f="2.371" m="0.726196" t="0.467"/>
// <eng-data cg="35." f="2.307" m="0.509957" t="0.532"/>
// <eng-data cg="35." f="2.371" m="0.320333" t="0.589"/>
// <eng-data cg="35." f="2.371" m="0.175326" t="0.632"/>
// <eng-data cg="35." f="2.243" m="0.109701" t="0.652"/>
// <eng-data cg="35." f="1.794" m="0.0637665" t="0.668"/>
// <eng-data cg="35." f="1.153" m="0.0302344" t="0.684"/>
// <eng-data cg="35." f="0.448" m="0.00860204" t="0.703"/>
// <eng-data cg="35." f="0." m="0." t="0.73"/>
// </data>
// </engine>
//  </engine-list>
// </engine-database>

func (r RSEParser) parse(filename string) *Parser{
	var home string = os.Getenv("HOME")
	var engines string = home+"/.gorocksim/engines/"+filename
	if _, err := os.Stat(engines); os.IsNotExist(err) {
		panic(err)
	}

	xmlFile, err := os.Open(engines)
	check(err)
	fmt.Println("Successfully Opened rse *.xml")
	defer xmlFile.Close()

	engine_info, err := ioutil.ReadAll(xmlFile)
	check(err)
	
	var engineDB EngineDatabase
	err = xml.Unmarshal(engine_info, &engineDB)
	if err != nil {
		fmt.Printf("xml.Unmarshal failed with '%s'\n", err)
	}
	the_engine := engineDB.Database.Engines[0]
	r.Context.Code = the_engine.Code
	r.Context.Diameter = the_engine.Diameter
	r.Context.Length = the_engine.Length
	r.Context.Delay = the_engine.Delay
	r.Context.PropWeight = the_engine.PropellantWeight
	r.Context.EngineWeight = the_engine.EngineWeight
	r.Context.Manufacturer = the_engine.Manufacturer
	r.Context.BurnTime = the_engine.BurnTime
	r.Context.Data = the_engine.Data
    
	return &r.Context
}

func (r RSEParser) setContext(parent EngineParser)  {
	r.Context = parent
}


func (r RSEParser) force_value_at(time float64) float64 {
	return 0.0
}



// doc = File.open(context.filename) { |f| Nokogiri::XML(f) }
// context.code = doc.xpath('string(//engine/@code)')
// context.diameter = doc.xpath('string(//engine/@dia)')
// context.length = doc.xpath('string(//engine/@len)')
// context.delay =doc.xpath('string(//engine/@delays)')
// context.propellant_weight = doc.xpath('number(//engine/@propWt)').to_f*0.001
// context.engine_weight =doc.xpath('number(//engine/@initWt)').to_f*0.001
// context.manufacturer  =doc.xpath('string(//engine/@mfg)')
// context.burn_time = doc.xpath('number(//engine/@burn-time)')
// doc.xpath('//eng-data').each do |node|
//   time_stamp = node.attribute('t').value.to_f
//   force = node.attribute('f').value.to_f
//   mass = ((context.engine_weight.to_f - context.propellant_weight) + node.attribute('m').value.to_f*0.001)
//   context.add_data_point([time_stamp, [force, mass]])
// end

//https://github.com/moovweb/gokogiri