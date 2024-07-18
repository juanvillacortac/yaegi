// Code generated by 'yaegi extract debug/macho'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"debug/macho"
	"reflect"
)

func init() {
	Symbols["debug/macho/macho"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ARM64_RELOC_ADDEND":              reflect.ValueOf(macho.ARM64_RELOC_ADDEND),
		"ARM64_RELOC_BRANCH26":            reflect.ValueOf(macho.ARM64_RELOC_BRANCH26),
		"ARM64_RELOC_GOT_LOAD_PAGE21":     reflect.ValueOf(macho.ARM64_RELOC_GOT_LOAD_PAGE21),
		"ARM64_RELOC_GOT_LOAD_PAGEOFF12":  reflect.ValueOf(macho.ARM64_RELOC_GOT_LOAD_PAGEOFF12),
		"ARM64_RELOC_PAGE21":              reflect.ValueOf(macho.ARM64_RELOC_PAGE21),
		"ARM64_RELOC_PAGEOFF12":           reflect.ValueOf(macho.ARM64_RELOC_PAGEOFF12),
		"ARM64_RELOC_POINTER_TO_GOT":      reflect.ValueOf(macho.ARM64_RELOC_POINTER_TO_GOT),
		"ARM64_RELOC_SUBTRACTOR":          reflect.ValueOf(macho.ARM64_RELOC_SUBTRACTOR),
		"ARM64_RELOC_TLVP_LOAD_PAGE21":    reflect.ValueOf(macho.ARM64_RELOC_TLVP_LOAD_PAGE21),
		"ARM64_RELOC_TLVP_LOAD_PAGEOFF12": reflect.ValueOf(macho.ARM64_RELOC_TLVP_LOAD_PAGEOFF12),
		"ARM64_RELOC_UNSIGNED":            reflect.ValueOf(macho.ARM64_RELOC_UNSIGNED),
		"ARM_RELOC_BR24":                  reflect.ValueOf(macho.ARM_RELOC_BR24),
		"ARM_RELOC_HALF":                  reflect.ValueOf(macho.ARM_RELOC_HALF),
		"ARM_RELOC_HALF_SECTDIFF":         reflect.ValueOf(macho.ARM_RELOC_HALF_SECTDIFF),
		"ARM_RELOC_LOCAL_SECTDIFF":        reflect.ValueOf(macho.ARM_RELOC_LOCAL_SECTDIFF),
		"ARM_RELOC_PAIR":                  reflect.ValueOf(macho.ARM_RELOC_PAIR),
		"ARM_RELOC_PB_LA_PTR":             reflect.ValueOf(macho.ARM_RELOC_PB_LA_PTR),
		"ARM_RELOC_SECTDIFF":              reflect.ValueOf(macho.ARM_RELOC_SECTDIFF),
		"ARM_RELOC_VANILLA":               reflect.ValueOf(macho.ARM_RELOC_VANILLA),
		"ARM_THUMB_32BIT_BRANCH":          reflect.ValueOf(macho.ARM_THUMB_32BIT_BRANCH),
		"ARM_THUMB_RELOC_BR22":            reflect.ValueOf(macho.ARM_THUMB_RELOC_BR22),
		"Cpu386":                          reflect.ValueOf(macho.Cpu386),
		"CpuAmd64":                        reflect.ValueOf(macho.CpuAmd64),
		"CpuArm":                          reflect.ValueOf(macho.CpuArm),
		"CpuArm64":                        reflect.ValueOf(macho.CpuArm64),
		"CpuPpc":                          reflect.ValueOf(macho.CpuPpc),
		"CpuPpc64":                        reflect.ValueOf(macho.CpuPpc64),
		"ErrNotFat":                       reflect.ValueOf(&macho.ErrNotFat).Elem(),
		"FlagAllModsBound":                reflect.ValueOf(macho.FlagAllModsBound),
		"FlagAllowStackExecution":         reflect.ValueOf(macho.FlagAllowStackExecution),
		"FlagAppExtensionSafe":            reflect.ValueOf(macho.FlagAppExtensionSafe),
		"FlagBindAtLoad":                  reflect.ValueOf(macho.FlagBindAtLoad),
		"FlagBindsToWeak":                 reflect.ValueOf(macho.FlagBindsToWeak),
		"FlagCanonical":                   reflect.ValueOf(macho.FlagCanonical),
		"FlagDeadStrippableDylib":         reflect.ValueOf(macho.FlagDeadStrippableDylib),
		"FlagDyldLink":                    reflect.ValueOf(macho.FlagDyldLink),
		"FlagForceFlat":                   reflect.ValueOf(macho.FlagForceFlat),
		"FlagHasTLVDescriptors":           reflect.ValueOf(macho.FlagHasTLVDescriptors),
		"FlagIncrLink":                    reflect.ValueOf(macho.FlagIncrLink),
		"FlagLazyInit":                    reflect.ValueOf(macho.FlagLazyInit),
		"FlagNoFixPrebinding":             reflect.ValueOf(macho.FlagNoFixPrebinding),
		"FlagNoHeapExecution":             reflect.ValueOf(macho.FlagNoHeapExecution),
		"FlagNoMultiDefs":                 reflect.ValueOf(macho.FlagNoMultiDefs),
		"FlagNoReexportedDylibs":          reflect.ValueOf(macho.FlagNoReexportedDylibs),
		"FlagNoUndefs":                    reflect.ValueOf(macho.FlagNoUndefs),
		"FlagPIE":                         reflect.ValueOf(macho.FlagPIE),
		"FlagPrebindable":                 reflect.ValueOf(macho.FlagPrebindable),
		"FlagPrebound":                    reflect.ValueOf(macho.FlagPrebound),
		"FlagRootSafe":                    reflect.ValueOf(macho.FlagRootSafe),
		"FlagSetuidSafe":                  reflect.ValueOf(macho.FlagSetuidSafe),
		"FlagSplitSegs":                   reflect.ValueOf(macho.FlagSplitSegs),
		"FlagSubsectionsViaSymbols":       reflect.ValueOf(macho.FlagSubsectionsViaSymbols),
		"FlagTwoLevel":                    reflect.ValueOf(macho.FlagTwoLevel),
		"FlagWeakDefines":                 reflect.ValueOf(macho.FlagWeakDefines),
		"GENERIC_RELOC_LOCAL_SECTDIFF":    reflect.ValueOf(macho.GENERIC_RELOC_LOCAL_SECTDIFF),
		"GENERIC_RELOC_PAIR":              reflect.ValueOf(macho.GENERIC_RELOC_PAIR),
		"GENERIC_RELOC_PB_LA_PTR":         reflect.ValueOf(macho.GENERIC_RELOC_PB_LA_PTR),
		"GENERIC_RELOC_SECTDIFF":          reflect.ValueOf(macho.GENERIC_RELOC_SECTDIFF),
		"GENERIC_RELOC_TLV":               reflect.ValueOf(macho.GENERIC_RELOC_TLV),
		"GENERIC_RELOC_VANILLA":           reflect.ValueOf(macho.GENERIC_RELOC_VANILLA),
		"LoadCmdDylib":                    reflect.ValueOf(macho.LoadCmdDylib),
		"LoadCmdDylinker":                 reflect.ValueOf(macho.LoadCmdDylinker),
		"LoadCmdDysymtab":                 reflect.ValueOf(macho.LoadCmdDysymtab),
		"LoadCmdRpath":                    reflect.ValueOf(macho.LoadCmdRpath),
		"LoadCmdSegment":                  reflect.ValueOf(macho.LoadCmdSegment),
		"LoadCmdSegment64":                reflect.ValueOf(macho.LoadCmdSegment64),
		"LoadCmdSymtab":                   reflect.ValueOf(macho.LoadCmdSymtab),
		"LoadCmdThread":                   reflect.ValueOf(macho.LoadCmdThread),
		"LoadCmdUnixThread":               reflect.ValueOf(macho.LoadCmdUnixThread),
		"Magic32":                         reflect.ValueOf(macho.Magic32),
		"Magic64":                         reflect.ValueOf(macho.Magic64),
		"MagicFat":                        reflect.ValueOf(macho.MagicFat),
		"NewFatFile":                      reflect.ValueOf(macho.NewFatFile),
		"NewFile":                         reflect.ValueOf(macho.NewFile),
		"Open":                            reflect.ValueOf(macho.Open),
		"OpenFat":                         reflect.ValueOf(macho.OpenFat),
		"TypeBundle":                      reflect.ValueOf(macho.TypeBundle),
		"TypeDylib":                       reflect.ValueOf(macho.TypeDylib),
		"TypeExec":                        reflect.ValueOf(macho.TypeExec),
		"TypeObj":                         reflect.ValueOf(macho.TypeObj),
		"X86_64_RELOC_BRANCH":             reflect.ValueOf(macho.X86_64_RELOC_BRANCH),
		"X86_64_RELOC_GOT":                reflect.ValueOf(macho.X86_64_RELOC_GOT),
		"X86_64_RELOC_GOT_LOAD":           reflect.ValueOf(macho.X86_64_RELOC_GOT_LOAD),
		"X86_64_RELOC_SIGNED":             reflect.ValueOf(macho.X86_64_RELOC_SIGNED),
		"X86_64_RELOC_SIGNED_1":           reflect.ValueOf(macho.X86_64_RELOC_SIGNED_1),
		"X86_64_RELOC_SIGNED_2":           reflect.ValueOf(macho.X86_64_RELOC_SIGNED_2),
		"X86_64_RELOC_SIGNED_4":           reflect.ValueOf(macho.X86_64_RELOC_SIGNED_4),
		"X86_64_RELOC_SUBTRACTOR":         reflect.ValueOf(macho.X86_64_RELOC_SUBTRACTOR),
		"X86_64_RELOC_TLV":                reflect.ValueOf(macho.X86_64_RELOC_TLV),
		"X86_64_RELOC_UNSIGNED":           reflect.ValueOf(macho.X86_64_RELOC_UNSIGNED),

		// type definitions
		"Cpu":              reflect.ValueOf((*macho.Cpu)(nil)),
		"Dylib":            reflect.ValueOf((*macho.Dylib)(nil)),
		"DylibCmd":         reflect.ValueOf((*macho.DylibCmd)(nil)),
		"Dysymtab":         reflect.ValueOf((*macho.Dysymtab)(nil)),
		"DysymtabCmd":      reflect.ValueOf((*macho.DysymtabCmd)(nil)),
		"FatArch":          reflect.ValueOf((*macho.FatArch)(nil)),
		"FatArchHeader":    reflect.ValueOf((*macho.FatArchHeader)(nil)),
		"FatFile":          reflect.ValueOf((*macho.FatFile)(nil)),
		"File":             reflect.ValueOf((*macho.File)(nil)),
		"FileHeader":       reflect.ValueOf((*macho.FileHeader)(nil)),
		"FormatError":      reflect.ValueOf((*macho.FormatError)(nil)),
		"Load":             reflect.ValueOf((*macho.Load)(nil)),
		"LoadBytes":        reflect.ValueOf((*macho.LoadBytes)(nil)),
		"LoadCmd":          reflect.ValueOf((*macho.LoadCmd)(nil)),
		"Nlist32":          reflect.ValueOf((*macho.Nlist32)(nil)),
		"Nlist64":          reflect.ValueOf((*macho.Nlist64)(nil)),
		"Regs386":          reflect.ValueOf((*macho.Regs386)(nil)),
		"RegsAMD64":        reflect.ValueOf((*macho.RegsAMD64)(nil)),
		"Reloc":            reflect.ValueOf((*macho.Reloc)(nil)),
		"RelocTypeARM":     reflect.ValueOf((*macho.RelocTypeARM)(nil)),
		"RelocTypeARM64":   reflect.ValueOf((*macho.RelocTypeARM64)(nil)),
		"RelocTypeGeneric": reflect.ValueOf((*macho.RelocTypeGeneric)(nil)),
		"RelocTypeX86_64":  reflect.ValueOf((*macho.RelocTypeX86_64)(nil)),
		"Rpath":            reflect.ValueOf((*macho.Rpath)(nil)),
		"RpathCmd":         reflect.ValueOf((*macho.RpathCmd)(nil)),
		"Section":          reflect.ValueOf((*macho.Section)(nil)),
		"Section32":        reflect.ValueOf((*macho.Section32)(nil)),
		"Section64":        reflect.ValueOf((*macho.Section64)(nil)),
		"SectionHeader":    reflect.ValueOf((*macho.SectionHeader)(nil)),
		"Segment":          reflect.ValueOf((*macho.Segment)(nil)),
		"Segment32":        reflect.ValueOf((*macho.Segment32)(nil)),
		"Segment64":        reflect.ValueOf((*macho.Segment64)(nil)),
		"SegmentHeader":    reflect.ValueOf((*macho.SegmentHeader)(nil)),
		"Symbol":           reflect.ValueOf((*macho.Symbol)(nil)),
		"Symtab":           reflect.ValueOf((*macho.Symtab)(nil)),
		"SymtabCmd":        reflect.ValueOf((*macho.SymtabCmd)(nil)),
		"Thread":           reflect.ValueOf((*macho.Thread)(nil)),
		"Type":             reflect.ValueOf((*macho.Type)(nil)),

		// interface wrapper definitions
		"_Load": reflect.ValueOf((*_debug_macho_Load)(nil)),
	}
}

// _debug_macho_Load is an interface wrapper for Load type
type _debug_macho_Load struct {
	IValue interface{}
	WRaw   func() []byte
}

func (W _debug_macho_Load) Raw() []byte { return W.WRaw() }
