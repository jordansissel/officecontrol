package main

type DenonChidleyRoot314159 struct {
	DenonItem *DenonItem `xml:" item,omitempty" json:"item,omitempty"`
}

type DenonItem struct {
	DenonAddSourceDisplay  *DenonAddSourceDisplay  `xml:" AddSourceDisplay,omitempty" json:"AddSourceDisplay,omitempty"`
	DenonBrandId           *DenonBrandId           `xml:" BrandId,omitempty" json:"BrandId,omitempty"`
	DenonECOMode           *DenonECOMode           `xml:" ECOMode,omitempty" json:"ECOMode,omitempty"`
	DenonECOModeDisp       *DenonECOModeDisp       `xml:" ECOModeDisp,omitempty" json:"ECOModeDisp,omitempty"`
	DenonECOModeLists      *DenonECOModeLists      `xml:" ECOModeLists,omitempty" json:"ECOModeLists,omitempty"`
	DenonFriendlyName      *DenonFriendlyName      `xml:" FriendlyName,omitempty" json:"FriendlyName,omitempty"`
	DenonInputFuncSelect   *DenonInputFuncSelect   `xml:" InputFuncSelect,omitempty" json:"InputFuncSelect,omitempty"`
	DenonMasterVolume      *DenonMasterVolume      `xml:" MasterVolume,omitempty" json:"MasterVolume,omitempty"`
	DenonModelId           *DenonModelId           `xml:" ModelId,omitempty" json:"ModelId,omitempty"`
	DenonMute              *DenonMute              `xml:" Mute,omitempty" json:"Mute,omitempty"`
	DenonNetFuncSelect     *DenonNetFuncSelect     `xml:" NetFuncSelect,omitempty" json:"NetFuncSelect,omitempty"`
	DenonPower             *DenonPower             `xml:" Power,omitempty" json:"Power,omitempty"`
	DenonRCSourceSelect    *DenonRCSourceSelect    `xml:" RCSourceSelect,omitempty" json:"RCSourceSelect,omitempty"`
	DenonRemoteMaintenance *DenonRemoteMaintenance `xml:" RemoteMaintenance,omitempty" json:"RemoteMaintenance,omitempty"`
	DenonRenameZone        *DenonRenameZone        `xml:" RenameZone,omitempty" json:"RenameZone,omitempty"`
	DenonSalesArea         *DenonSalesArea         `xml:" SalesArea,omitempty" json:"SalesArea,omitempty"`
	DenonSelectSurround    *DenonSelectSurround    `xml:" selectSurround,omitempty" json:"selectSurround,omitempty"`
	DenonSleepOff          *DenonSleepOff          `xml:" SleepOff,omitempty" json:"SleepOff,omitempty"`
	DenonSubwooferDisplay  *DenonSubwooferDisplay  `xml:" SubwooferDisplay,omitempty" json:"SubwooferDisplay,omitempty"`
	DenonTopMenuLink       *DenonTopMenuLink       `xml:" TopMenuLink,omitempty" json:"TopMenuLink,omitempty"`
	DenonVideoSelect       *DenonVideoSelect       `xml:" VideoSelect,omitempty" json:"VideoSelect,omitempty"`
	DenonVideoSelectDisp   *DenonVideoSelectDisp   `xml:" VideoSelectDisp,omitempty" json:"VideoSelectDisp,omitempty"`
	DenonVideoSelectLists  *DenonVideoSelectLists  `xml:" VideoSelectLists,omitempty" json:"VideoSelectLists,omitempty"`
	DenonVideoSelectOnOff  *DenonVideoSelectOnOff  `xml:" VideoSelectOnOff,omitempty" json:"VideoSelectOnOff,omitempty"`
	DenonVolumeDisplay     *DenonVolumeDisplay     `xml:" VolumeDisplay,omitempty" json:"VolumeDisplay,omitempty"`
	DenonZone2VolDisp      *DenonZone2VolDisp      `xml:" Zone2VolDisp,omitempty" json:"Zone2VolDisp,omitempty"`
	DenonZonePower         *DenonZonePower         `xml:" ZonePower,omitempty" json:"ZonePower,omitempty"`
}

type DenonFriendlyName struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonValue struct {
	AttrIndex string `xml:" index,attr"  json:",omitempty"`
	AttrParam string `xml:" param,attr"  json:",omitempty"`
	AttrTable string `xml:" table,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type DenonPower struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonZonePower struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonRCSourceSelect struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonRenameZone struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonTopMenuLink struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonVideoSelectDisp struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonVideoSelect struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonVideoSelectOnOff struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonVideoSelectLists struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonECOModeDisp struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonECOMode struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonECOModeLists struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonAddSourceDisplay struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonModelId struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonBrandId struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonSalesArea struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonInputFuncSelect struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonNetFuncSelect struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonSelectSurround struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonVolumeDisplay struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonMasterVolume struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonMute struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonRemoteMaintenance struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonSubwooferDisplay struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonZone2VolDisp struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}

type DenonSleepOff struct {
	DenonValue []*DenonValue `xml:" value,omitempty" json:"value,omitempty"`
}
