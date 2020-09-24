package drda

const (
	DrdaMagic         uint16 = 0xD0
	DrdaCpData        uint16 = 0x0000
	DrdaCpCodpnt      uint16 = 0x000C
	DrdaCpFdodsc      uint16 = 0x0010
	DrdaCpTypdefnam   uint16 = 0x002F
	DrdaCpTypdefovr   uint16 = 0x0035
	DrdaCpCodpntdr    uint16 = 0x0064
	DrdaCpExcsat      uint16 = 0x1041
	DrdaCpSyncctl     uint16 = 0x1055
	DrdaCpSyncrsy     uint16 = 0x1069
	DrdaCpAccsec      uint16 = 0x106D
	DrdaCpSecchk      uint16 = 0x106E
	DrdaCpSynclog     uint16 = 0x106F
	DrdaCpRsctyp      uint16 = 0x111F
	DrdaCpRsncod      uint16 = 0x1127
	DrdaCpRscnam      uint16 = 0x112D
	DrdaCpPrdid       uint16 = 0x112E
	DrdaCpPrccnvcd    uint16 = 0x113F
	DrdaCpVrsnam      uint16 = 0x1144
	DrdaCpSrvclsnm    uint16 = 0x1147
	DrdaCpSvrcod      uint16 = 0x1149
	DrdaCpSynerrcd    uint16 = 0x114A
	DrdaCpSrvdgn      uint16 = 0x1153
	DrdaCpSrvrlslv    uint16 = 0x115A
	DrdaCpSpvnam      uint16 = 0x115D
	DrdaCpExtnam      uint16 = 0x115E
	DrdaCpSrvnam      uint16 = 0x116D
	DrdaCpSecmgrnm    uint16 = 0x1196
	DrdaCpDeperrcd    uint16 = 0x119B
	DrdaCpCcsidsbc    uint16 = 0x119C
	DrdaCpCcsiddbc    uint16 = 0x119D
	DrdaCpCcsidmbc    uint16 = 0x119E
	DrdaCpUsrid       uint16 = 0x11A0
	DrdaCpPassword    uint16 = 0x11A1
	DrdaCpSecmec      uint16 = 0x11A2
	DrdaCpSecchkcd    uint16 = 0x11A4
	DrdaCpSvcerrno    uint16 = 0x11B4
	DrdaCpSectkn      uint16 = 0x11DC
	DrdaCpNewpassword uint16 = 0x11DE
	DrdaCpMgrlvlrm    uint16 = 0x1210
	DrdaCpMgrdeprm    uint16 = 0x1218
	DrdaCpSecchkrm    uint16 = 0x1219
	DrdaCpCmdathrm    uint16 = 0x121C
	DrdaCpAgnprmrm    uint16 = 0x1232
	DrdaCpRsclmtrm    uint16 = 0x1233
	DrdaCpPrccnvrm    uint16 = 0x1245
	DrdaCpCmdcmprm    uint16 = 0x124B
	DrdaCpSyntaxrm    uint16 = 0x124C
	DrdaCpCmdnsprm    uint16 = 0x1250
	DrdaCpPrmnsprm    uint16 = 0x1251
	DrdaCpValnsprm    uint16 = 0x1252
	DrdaCpObjnsprm    uint16 = 0x1253
	DrdaCpCmdchkrm    uint16 = 0x1254
	DrdaCpTrgnsprm    uint16 = 0x125F
	DrdaCpAgent       uint16 = 0x1403
	DrdaCpMgrlvlls    uint16 = 0x1404
	DrdaCpSupervisor  uint16 = 0x143C
	DrdaCpSecmgr      uint16 = 0x1440
	DrdaCpExcsatrd    uint16 = 0x1443
	DrdaCpCmnappc     uint16 = 0x1444
	DrdaCpDictionary  uint16 = 0x1458
	DrdaCpMgrlvln     uint16 = 0x1473
	DrdaCpCmntcpip    uint16 = 0x1474
	DrdaCpFdodta      uint16 = 0x147A
	DrdaCpCmnsyncpt   uint16 = 0x147C
	DrdaCpAccsecrd    uint16 = 0x14AC
	DrdaCpSyncptmgr   uint16 = 0x14C0
	DrdaCpRsyncmgr    uint16 = 0x14C1
	DrdaCpCcsidmgr    uint16 = 0x14CC
	DrdaCpMonitor     uint16 = 0x1900
	DrdaCpMonitorrd   uint16 = 0x1C00
	DrdaCpXamgr       uint16 = 0x1C01
	DrdaCpAccrdb      uint16 = 0x2001
	DrdaCpBgnbnd      uint16 = 0x2002
	DrdaCpBndsqlstt   uint16 = 0x2004
	DrdaCpClsqry      uint16 = 0x2005
	DrdaCpCntqry      uint16 = 0x2006
	DrdaCpDrppkg      uint16 = 0x2007
	DrdaCpDscsqlstt   uint16 = 0x2008
	DrdaCpEndbnd      uint16 = 0x2009
	DrdaCpExcsqlimm   uint16 = 0x200A
	DrdaCpExcsqlstt   uint16 = 0x200B
	DrdaCpOpnqry      uint16 = 0x200C
	DrdaCpPrpsqlstt   uint16 = 0x200D
	DrdaCpRdbcmm      uint16 = 0x200E
	DrdaCpRdbrllbck   uint16 = 0x200F
	DrdaCpRebind      uint16 = 0x2010
	DrdaCpDscrdbtbl   uint16 = 0x2012
	DrdaCpExcsqlset   uint16 = 0x2014
	DrdaCpDscerrcd    uint16 = 0x2101
	DrdaCpQryprctyp   uint16 = 0x2102
	DrdaCpRdbinttkn   uint16 = 0x2103
	DrdaCpPrddta      uint16 = 0x2104
	DrdaCpRdbcmtok    uint16 = 0x2105
	DrdaCpRdbcolid    uint16 = 0x2108
	DrdaCpPkgid       uint16 = 0x2109
	DrdaCpPkgcnstkn   uint16 = 0x210D
	DrdaCpRtnsetstt   uint16 = 0x210E
	DrdaCpRdbacccl    uint16 = 0x210F
	DrdaCpRdbnam      uint16 = 0x2110
	DrdaCpOutexp      uint16 = 0x2111
	DrdaCpPkgnamct    uint16 = 0x2112
	DrdaCpPkgnamcsn   uint16 = 0x2113
	DrdaCpQryblksz    uint16 = 0x2114
	DrdaCpUowdsp      uint16 = 0x2115
	DrdaCpRtnsqlda    uint16 = 0x2116
	DrdaCpRdbalwupd   uint16 = 0x211A
	DrdaCpSqlcsrhld   uint16 = 0x211F
	DrdaCpSttstrdel   uint16 = 0x2120
	DrdaCpSttdecdel   uint16 = 0x2121
	DrdaCpPkgdftcst   uint16 = 0x2125
	DrdaCpQryblkctl   uint16 = 0x2132
	DrdaCpCrrtkn      uint16 = 0x2135
	DrdaCpPrcnam      uint16 = 0x2138
	DrdaCpPkgsnlst    uint16 = 0x2139
	DrdaCpNbrrow      uint16 = 0x213A
	DrdaCpTrgdftrt    uint16 = 0x213B
	DrdaCpQryrelscr   uint16 = 0x213C
	DrdaCpQryrownbr   uint16 = 0x213D
	DrdaCpQryrfrtbl   uint16 = 0x213E
	DrdaCpMaxrslcnt   uint16 = 0x2140
	DrdaCpMaxblkext   uint16 = 0x2141
	DrdaCpRslsetflg   uint16 = 0x2142
	DrdaCpTypsqlda    uint16 = 0x2146
	DrdaCpOutovropt   uint16 = 0x2147
	DrdaCpRtnextdta   uint16 = 0x2148
	DrdaCpQryattscr   uint16 = 0x2149
	DrdaCpQryattupd   uint16 = 0x2150
	DrdaCpQryscrorn   uint16 = 0x2152
	DrdaCpQryrowsns   uint16 = 0x2153
	DrdaCpQryblkrst   uint16 = 0x2154
	DrdaCpQryrtndta   uint16 = 0x2155
	DrdaCpQryrowset   uint16 = 0x2156
	DrdaCpQryattsns   uint16 = 0x2157
	DrdaCpQryinsid    uint16 = 0x215B
	DrdaCpQryclsimp   uint16 = 0x215D
	DrdaCpQryclsrls   uint16 = 0x215E
	DrdaCpQryoptval   uint16 = 0x215F
	DrdaCpDiaglvl     uint16 = 0x2160
	DrdaCpAccrdbrm    uint16 = 0x2201
	DrdaCpQrynoprm    uint16 = 0x2202
	DrdaCpRdbnacrm    uint16 = 0x2204
	DrdaCpOpnqryrm    uint16 = 0x2205
	DrdaCpPkgbnarm    uint16 = 0x2206
	DrdaCpRdbaccrm    uint16 = 0x2207
	DrdaCpBgnbndrm    uint16 = 0x2208
	DrdaCpPkgbparm    uint16 = 0x2209
	DrdaCpDscinvrm    uint16 = 0x220A
	DrdaCpEndqryrm    uint16 = 0x220B
	DrdaCpEnduowrm    uint16 = 0x220C
	DrdaCpAbnuowrm    uint16 = 0x220D
	DrdaCpDtamchrm    uint16 = 0x220E
	DrdaCpQrypoprm    uint16 = 0x220F
	DrdaCpRdbnfnrm    uint16 = 0x2211
	DrdaCpOpnqflrm    uint16 = 0x2212
	DrdaCpSqlerrrm    uint16 = 0x2213
	DrdaCpRdbupdrm    uint16 = 0x2218
	DrdaCpRslsetrm    uint16 = 0x2219
	DrdaCpRdbaflrm    uint16 = 0x221A
	DrdaCpCmdvltrm    uint16 = 0x221D
	DrdaCpCmmrqsrm    uint16 = 0x2225
	DrdaCpRdbathrm    uint16 = 0x22CB
	DrdaCpSqlam       uint16 = 0x2407
	DrdaCpSqlcard     uint16 = 0x2408
	DrdaCpSqlcinrd    uint16 = 0x240B
	DrdaCpSqlrslrd    uint16 = 0x240E
	DrdaCpRdb         uint16 = 0x240F
	DrdaCpFrcfixrow   uint16 = 0x2410
	DrdaCpSqldard     uint16 = 0x2411
	DrdaCpSqldta      uint16 = 0x2412
	DrdaCpSqldtard    uint16 = 0x2413
	DrdaCpSqlstt      uint16 = 0x2414
	DrdaCpOutovr      uint16 = 0x2415
	DrdaCpLmtblkprc   uint16 = 0x2417
	DrdaCpFixrowprc   uint16 = 0x2418
	DrdaCpSqlsttvrb   uint16 = 0x2419
	DrdaCpQrydsc      uint16 = 0x241A
	DrdaCpQrydta      uint16 = 0x241B
	DrdaCpCstmbcs     uint16 = 0x2435
	DrdaCpSrvlst      uint16 = 0x244E
	DrdaCpSqlattr     uint16 = 0x2450

	// --- Product-specific 0xC000-0xFFFF ---
	// Piggy-backed session data (product-specific)
	DrdaCpPbsd uint16 = 0xC000

	// Isolation level as a byte (product-specific)
	DrdaCpPbsdIso uint16 = 0xC001

	// Current schema as UTF8 String (product-specific)
	DrdaCpPbsdSchema uint16 = 0xC002

	DrdaDssfmtSameCorr uint16 = 0x01
	DrdaDssfmtContinue uint16 = 0x02
	DrdaDssfmtChained  uint16 = 0x04
	DrdaDssfmtReserved uint16 = 0x08

	DrdaDssfmtRqsdss   uint16 = 0x01
	DrdaDssfmtRpydss   uint16 = 0x02
	DrdaDssfmtObjdss   uint16 = 0x03
	DrdaDssfmtCmndss   uint16 = 0x04
	DrdaDssfmtNorpydss uint16 = 0x05

	DrdaTextDdm   = "DDM"
	DrdaTextParam = "Parameter"
)

var DRDA_Description = map[uint16]string{
	DrdaCpData:        "Data",
	DrdaCpCodpnt:      "Code Point",
	DrdaCpFdodsc:      "FD:OCA Data Descriptor",
	DrdaCpTypdefnam:   "Data Type Definition Name",
	DrdaCpTypdefovr:   "TYPDEF Overrides",
	DrdaCpCodpntdr:    "Code Point Data Representation",
	DrdaCpExcsat:      "Exchange Server Attributes",
	DrdaCpSyncctl:     "Sync Point Control Request",
	DrdaCpSyncrsy:     "Sync Point Resync Command",
	DrdaCpAccsec:      "Access Security",
	DrdaCpSecchk:      "Security Check",
	DrdaCpSynclog:     "Sync Point Log",
	DrdaCpRsctyp:      "Resource Type Information",
	DrdaCpRsncod:      "Reason Code Information",
	DrdaCpRscnam:      "Resource Name Information",
	DrdaCpPrdid:       "Product-Specific Identifier",
	DrdaCpPrccnvcd:    "Conversation Protocol Error Code",
	DrdaCpVrsnam:      "Version Name",
	DrdaCpSrvclsnm:    "Server Class Name",
	DrdaCpSvrcod:      "Severity Code",
	DrdaCpSynerrcd:    "Syntax Error Code",
	DrdaCpSrvdgn:      "Server Diagnostic Information",
	DrdaCpSrvrlslv:    "Server Product Release Level",
	DrdaCpSpvnam:      "Supervisor Name",
	DrdaCpExtnam:      "External Name",
	DrdaCpSrvnam:      "Server Name",
	DrdaCpSecmgrnm:    "Security Manager Name",
	DrdaCpDeperrcd:    "Manager Dependency Error Code",
	DrdaCpCcsidsbc:    "CCSID for Single-Byte Characters",
	DrdaCpCcsiddbc:    "CCSID for Double-byte Characters",
	DrdaCpCcsidmbc:    "CCSID for Mixed-byte Characters",
	DrdaCpUsrid:       "User ID at the Target System",
	DrdaCpPassword:    "Password",
	DrdaCpSecmec:      "Security Mechanism",
	DrdaCpSecchkcd:    "Security Check Code",
	DrdaCpSvcerrno:    "Security Service ErrorNumber",
	DrdaCpSectkn:      "Security Token",
	DrdaCpNewpassword: "New Password",
	DrdaCpMgrlvlrm:    "Manager-Level Conflict",
	DrdaCpMgrdeprm:    "Manager Dependency Error",
	DrdaCpSecchkrm:    "Security Check",
	DrdaCpCmdathrm:    "Not Authorized to Command",
	DrdaCpAgnprmrm:    "Permanent Agent Error",
	DrdaCpRsclmtrm:    "Resource Limits Reached",
	DrdaCpPrccnvrm:    "Conversational Protocol Error",
	DrdaCpCmdcmprm:    "Command Processing Completed",
	DrdaCpSyntaxrm:    "Data Stream Syntax Error",
	DrdaCpCmdnsprm:    "Command Not Supported",
	DrdaCpPrmnsprm:    "Parameter Not Supported",
	DrdaCpValnsprm:    "Parameter Value Not Supported",
	DrdaCpObjnsprm:    "Object Not Supported",
	DrdaCpCmdchkrm:    "Command Check",
	DrdaCpTrgnsprm:    "Target Not Supported",
	DrdaCpAgent:       "Agent",
	DrdaCpMgrlvlls:    "Manager-Level List",
	DrdaCpSupervisor:  "Supervisor",
	DrdaCpSecmgr:      "Security Manager",
	DrdaCpExcsatrd:    "Server Attributes Reply Data",
	DrdaCpCmnappc:     "LU 6.2 Conversational Communications Manager",
	DrdaCpDictionary:  "Dictionary",
	DrdaCpMgrlvln:     "Manager-Level Number Attribute",
	DrdaCpCmntcpip:    "TCP/IP CommunicationManager",
	DrdaCpFdodta:      "FD:OCA Data",
	DrdaCpCmnsyncpt:   "SNA LU 6.2 Sync Point Conversational Communications Manager",
	DrdaCpAccsecrd:    "Access Security Reply Data",
	DrdaCpSyncptmgr:   "Sync Point Manager",
	DrdaCpRsyncmgr:    "ResynchronizationManager",
	DrdaCpCcsidmgr:    "CCSID Manager",
	DrdaCpMonitor:     "Monitor Events",
	DrdaCpMonitorrd:   "Monitor Reply Data",
	DrdaCpXamgr:       "XAManager",
	DrdaCpAccrdb:      "Access RDB",
	DrdaCpBgnbnd:      "Begin Binding a Package to an RDB",
	DrdaCpBndsqlstt:   "Bind SQL Statement to an RDB Package",
	DrdaCpClsqry:      "Close Query",
	DrdaCpCntqry:      "Continue Query",
	DrdaCpDrppkg:      "Drop RDB Package",
	DrdaCpDscsqlstt:   "Describe SQL Statement",
	DrdaCpEndbnd:      "End Binding a Package to an RDB",
	DrdaCpExcsqlimm:   "Execute Immediate SQL Statement",
	DrdaCpExcsqlstt:   "Execute SQL Statement",
	DrdaCpOpnqry:      "Open Query",
	DrdaCpPrpsqlstt:   "Prepare SQL Statement",
	DrdaCpRdbcmm:      "RDB Commit Unit of Work",
	DrdaCpRdbrllbck:   "RDB Rollback Unit of Work",
	DrdaCpRebind:      "Rebind an Existing RDB Package",
	DrdaCpDscrdbtbl:   "Describe RDB Table",
	DrdaCpExcsqlset:   "Set SQL Environment",
	DrdaCpDscerrcd:    "Description Error Code",
	DrdaCpQryprctyp:   "Query Protocol Type",
	DrdaCpRdbinttkn:   "RDB Interrupt Token",
	DrdaCpPrddta:      "Product-Specific Data",
	DrdaCpRdbcmtok:    "RDB Commit Allowed",
	DrdaCpRdbcolid:    "RDB Collection Identifier",
	DrdaCpPkgid:       "RDB Package Identifier",
	DrdaCpPkgcnstkn:   "RDB Package Consistency Token",
	DrdaCpRtnsetstt:   "Return SET Statement",
	DrdaCpRdbacccl:    "RDB Access Manager Class",
	DrdaCpRdbnam:      "Relational Database Name",
	DrdaCpOutexp:      "Output Expected",
	DrdaCpPkgnamct:    "RDB Package Name and Consistency Token",
	DrdaCpPkgnamcsn:   "RDB Package Name, Consistency Token, and Section Number",
	DrdaCpQryblksz:    "Query Block Size",
	DrdaCpUowdsp:      "Unit of Work Disposition",
	DrdaCpRtnsqlda:    "Maximum Result Set Count",
	DrdaCpRdbalwupd:   "RDB Allow Updates",
	DrdaCpSqlcsrhld:   "Hold Cursor Position",
	DrdaCpSttstrdel:   "Statement String Delimiter",
	DrdaCpSttdecdel:   "Statement Decimal Delimiter",
	DrdaCpPkgdftcst:   "Package Default Character Subtype",
	DrdaCpQryblkctl:   "Query Block Protocol Control",
	DrdaCpCrrtkn:      "Correlation Token",
	DrdaCpPrcnam:      "Procedure Name",
	DrdaCpPkgsnlst:    "RDB Result Set Reply Message",
	DrdaCpNbrrow:      "Number of Fetch or Insert Rows",
	DrdaCpTrgdftrt:    "Target Default Value Return",
	DrdaCpQryrelscr:   "Query Relative Scrolling Action",
	DrdaCpQryrownbr:   "Query Row Number",
	DrdaCpQryrfrtbl:   "Query Refresh Answer Set Table",
	DrdaCpMaxrslcnt:   "Maximum Result Set Count",
	DrdaCpMaxblkext:   "Maximum Number of Extra Blocks",
	DrdaCpRslsetflg:   "Result Set Flags",
	DrdaCpTypsqlda:    "Type of SQL Descriptor Area",
	DrdaCpOutovropt:   "Output Override Option",
	DrdaCpRtnextdta:   "Return of EXTDTA Option",
	DrdaCpQryattscr:   "Query Attribute for Scrollability",
	DrdaCpQryattupd:   "Query Attribute for Updatability",
	DrdaCpQryscrorn:   "Query Scroll Orientation",
	DrdaCpQryrowsns:   "Query Row Sensitivity",
	DrdaCpQryblkrst:   "Query Block Reset",
	DrdaCpQryrtndta:   "Query Returns Datat",
	DrdaCpQryrowset:   "Query Rowset Size",
	DrdaCpQryattsns:   "Query Attribute for Sensitivity",
	DrdaCpQryinsid:    "Query Instance Identifier",
	DrdaCpQryclsimp:   "Query Close Implicit",
	DrdaCpQryclsrls:   "Query Close Lock Release",
	DrdaCpQryoptval:   "QRYOPTVAL",
	DrdaCpDiaglvl:     "SQL Error Diagnostic Level",
	DrdaCpAccrdbrm:    "Access to RDB Completed",
	DrdaCpQrynoprm:    "Query Not Open",
	DrdaCpRdbnacrm:    "RDB Not Accessed",
	DrdaCpOpnqryrm:    "Open Query Complete",
	DrdaCpPkgbnarm:    "RDB Package Binding Not Active",
	DrdaCpRdbaccrm:    "RDB Currently Accessed",
	DrdaCpBgnbndrm:    "Begin Bind Error",
	DrdaCpPkgbparm:    "RDB Package Binding Process Active",
	DrdaCpDscinvrm:    "Invalid Description",
	DrdaCpEndqryrm:    "End of Query",
	DrdaCpEnduowrm:    "End Unit of Work Condition",
	DrdaCpAbnuowrm:    "Abnormal End Unit ofWork Condition",
	DrdaCpDtamchrm:    "Data Descriptor Mismatch",
	DrdaCpQrypoprm:    "Query Previously Opened",
	DrdaCpRdbnfnrm:    "RDB Not Found",
	DrdaCpOpnqflrm:    "Open Query Failure",
	DrdaCpSqlerrrm:    "SQL Error Condition",
	DrdaCpRdbupdrm:    "RDB Update Reply Message",
	DrdaCpRslsetrm:    "RDB Result Set Reply Message",
	DrdaCpRdbaflrm:    "RDB Access Failed Reply Message",
	DrdaCpCmdvltrm:    "Command Violation",
	DrdaCpCmmrqsrm:    "Commitment Request",
	DrdaCpRdbathrm:    "Not Authorized to RDB",
	DrdaCpSqlam:       "SQL Application Manager",
	DrdaCpSqlcard:     "SQL Communications Area Reply Data",
	DrdaCpSqlcinrd:    "SQL Result Set Column Information Reply Data",
	DrdaCpSqlrslrd:    "SQL Result Set Reply Data",
	DrdaCpRdb:         "Relational Database",
	DrdaCpFrcfixrow:   "Force Fixed Row Query Protocol",
	DrdaCpSqldard:     "SQLDA Reply Data",
	DrdaCpSqldta:      "SQL Program Variable Data",
	DrdaCpSqldtard:    "SQL Data Reply Data",
	DrdaCpSqlstt:      "SQL Statement",
	DrdaCpOutovr:      "Output Override Descriptor",
	DrdaCpLmtblkprc:   "Limited Block Protocol",
	DrdaCpFixrowprc:   "Fixed Row Query Protocol",
	DrdaCpSqlsttvrb:   "SQL Statement Variable Descriptions",
	DrdaCpQrydsc:      "Query Answer Set Description",
	DrdaCpQrydta:      "Query Answer Set Data",
	DrdaCpSqlattr:     "SQL Statement Attributes",
	// Piggy-backed session data (product-specific)
	DrdaCpPbsd: "Piggy-backed session data (product-specific)",

	// Isolation level as a byte (product-specific)
	DrdaCpPbsdIso: "Isolation level as a byte (product-specific)",

	// Current schema as UTF8 String (product-specific)
	DrdaCpPbsdSchema: "Current schema as UTF8 String (product-specific)",
}

var DRDA_Abbrev = map[uint16]string{
	DrdaCpData:        "DATA",
	DrdaCpCodpnt:      "CODPNT",
	DrdaCpFdodsc:      "FDODSC",
	DrdaCpTypdefnam:   "TYPDEFNAM",
	DrdaCpTypdefovr:   "TYPDEFOVR",
	DrdaCpCodpntdr:    "CODPNTDR",
	DrdaCpExcsat:      "EXCSAT",
	DrdaCpSyncctl:     "SYNCCTL",
	DrdaCpSyncrsy:     "SYNCRSY",
	DrdaCpAccsec:      "ACCSEC",
	DrdaCpSecchk:      "SECCHK",
	DrdaCpSynclog:     "SYNCLOG",
	DrdaCpRsctyp:      "RSCTYP",
	DrdaCpRsncod:      "RSNCOD",
	DrdaCpRscnam:      "RSCNAM",
	DrdaCpPrdid:       "PRDID",
	DrdaCpPrccnvcd:    "PRCCNVCD",
	DrdaCpVrsnam:      "VRSNAM",
	DrdaCpSrvclsnm:    "SRVCLSNM",
	DrdaCpSvrcod:      "SVRCOD",
	DrdaCpSynerrcd:    "SYNERRCD",
	DrdaCpSrvdgn:      "SRVDGN",
	DrdaCpSrvrlslv:    "SRVRLSLV",
	DrdaCpSpvnam:      "SPVNAM",
	DrdaCpExtnam:      "EXTNAM",
	DrdaCpSrvnam:      "SRVNAM",
	DrdaCpSecmgrnm:    "SECMGRNM",
	DrdaCpDeperrcd:    "DEPERRCD",
	DrdaCpCcsidsbc:    "CCSIDSBC",
	DrdaCpCcsiddbc:    "CCSIDDBC",
	DrdaCpCcsidmbc:    "CCSIDMBC",
	DrdaCpUsrid:       "USRID",
	DrdaCpPassword:    "PASSWORD",
	DrdaCpSecmec:      "SECMEC",
	DrdaCpSecchkcd:    "SECCHKCD",
	DrdaCpSvcerrno:    "SVCERRNO",
	DrdaCpSectkn:      "SECTKN",
	DrdaCpNewpassword: "NEWPASSWORD",
	DrdaCpMgrlvlrm:    "MGRLVLRM",
	DrdaCpMgrdeprm:    "MGRDEPRM",
	DrdaCpSecchkrm:    "SECCHKRM",
	DrdaCpCmdathrm:    "CMDATHRM",
	DrdaCpAgnprmrm:    "AGNPRMRM",
	DrdaCpRsclmtrm:    "RSCLMTRM",
	DrdaCpPrccnvrm:    "PRCCNVRM",
	DrdaCpCmdcmprm:    "CMDCMPRM",
	DrdaCpSyntaxrm:    "SYNTAXRM",
	DrdaCpCmdnsprm:    "CMDNSPRM",
	DrdaCpPrmnsprm:    "PRMNSPRM",
	DrdaCpValnsprm:    "VALNSPRM",
	DrdaCpObjnsprm:    "OBJNSPRM",
	DrdaCpCmdchkrm:    "CMDCHKRM",
	DrdaCpTrgnsprm:    "TRGNSPRM",
	DrdaCpAgent:       "AGENT",
	DrdaCpMgrlvlls:    "MGRLVLLS",
	DrdaCpSupervisor:  "SUPERVISOR",
	DrdaCpSecmgr:      "SECMGR",
	DrdaCpExcsatrd:    "EXCSATRD",
	DrdaCpCmnappc:     "CMNAPPC",
	DrdaCpDictionary:  "DICTIONARY",
	DrdaCpMgrlvln:     "MGRLVLN",
	DrdaCpCmntcpip:    "CMNTCPIP",
	DrdaCpFdodta:      "FDODTA",
	DrdaCpCmnsyncpt:   "CMNSYNCPT",
	DrdaCpAccsecrd:    "ACCSECRD",
	DrdaCpSyncptmgr:   "SYNCPTMGR",
	DrdaCpRsyncmgr:    "RSYNCMGR",
	DrdaCpCcsidmgr:    "CCSIDMGR",
	DrdaCpMonitor:     "MONITOR",
	DrdaCpMonitorrd:   "MONITORRD",
	DrdaCpXamgr:       "XAMGR",
	DrdaCpAccrdb:      "ACCRDB",
	DrdaCpBgnbnd:      "BGNBND",
	DrdaCpBndsqlstt:   "BNDSQLSTT",
	DrdaCpClsqry:      "CLSQRY",
	DrdaCpCntqry:      "CNTQRY",
	DrdaCpDrppkg:      "DRPPKG",
	DrdaCpDscsqlstt:   "DSCSQLSTT",
	DrdaCpEndbnd:      "ENDBND",
	DrdaCpExcsqlimm:   "EXCSQLIMM",
	DrdaCpExcsqlstt:   "EXCSQLSTT",
	DrdaCpOpnqry:      "OPNQRY",
	DrdaCpPrpsqlstt:   "PRPSQLSTT",
	DrdaCpRdbcmm:      "RDBCMM",
	DrdaCpRdbrllbck:   "RDBRLLBCK",
	DrdaCpRebind:      "REBIND",
	DrdaCpDscrdbtbl:   "DSCRDBTBL",
	DrdaCpExcsqlset:   "EXCSQLSET",
	DrdaCpDscerrcd:    "DSCERRCD",
	DrdaCpQryprctyp:   "QRYPRCTYP",
	DrdaCpRdbinttkn:   "RDBINTTKN",
	DrdaCpPrddta:      "PRDDTA",
	DrdaCpRdbcmtok:    "RDBCMTOK",
	DrdaCpRdbcolid:    "RDBCOLID",
	DrdaCpPkgid:       "PKGID",
	DrdaCpPkgcnstkn:   "PKGCNSTKN",
	DrdaCpRtnsetstt:   "RTNSETSTT",
	DrdaCpRdbacccl:    "RDBACCCL",
	DrdaCpRdbnam:      "RDBNAM",
	DrdaCpOutexp:      "OUTEXP",
	DrdaCpPkgnamct:    "PKGNAMCT",
	DrdaCpPkgnamcsn:   "PKGNAMCSN",
	DrdaCpQryblksz:    "QRYBLKSZ",
	DrdaCpUowdsp:      "UOWDSP",
	DrdaCpRtnsqlda:    "RTNSQLDA",
	DrdaCpRdbalwupd:   "RDBALWUPD",
	DrdaCpSqlcsrhld:   "SQLCSRHLD",
	DrdaCpSttstrdel:   "STTSTRDEL",
	DrdaCpSttdecdel:   "STTDECDEL",
	DrdaCpPkgdftcst:   "PKGDFTCST",
	DrdaCpQryblkctl:   "QRYBLKCTL",
	DrdaCpCrrtkn:      "CRRTKN",
	DrdaCpPrcnam:      "PRCNAM",
	DrdaCpPkgsnlst:    "PKGSNLST",
	DrdaCpNbrrow:      "NBRROW",
	DrdaCpTrgdftrt:    "TRGDFTRT",
	DrdaCpQryrelscr:   "QRYRELSCR",
	DrdaCpQryrownbr:   "QRYROWNBR",
	DrdaCpQryrfrtbl:   "QRYRFRTBL",
	DrdaCpMaxrslcnt:   "MAXRSLCNT",
	DrdaCpMaxblkext:   "MAXBLKEXT",
	DrdaCpRslsetflg:   "RSLSETFLG",
	DrdaCpTypsqlda:    "TYPSQLDA",
	DrdaCpOutovropt:   "OUTOVROPT",
	DrdaCpRtnextdta:   "RTNEXTDTA",
	DrdaCpQryattscr:   "QRYATTSCR",
	DrdaCpQryattupd:   "QRYATTUPD",
	DrdaCpQryscrorn:   "QRYSCRORN",
	DrdaCpQryrowsns:   "QRYROWSNS",
	DrdaCpQryblkrst:   "QRYBLKRST",
	DrdaCpQryrtndta:   "QRYRTNDTA",
	DrdaCpQryrowset:   "QRYROWSET",
	DrdaCpQryattsns:   "QRYATTSNS",
	DrdaCpQryinsid:    "QRYINSID",
	DrdaCpQryclsimp:   "QRYCLSIMP",
	DrdaCpQryclsrls:   "QRYCLSRLS",
	DrdaCpQryoptval:   "QRYOPTVAL",
	DrdaCpDiaglvl:     "DIAGLVL",
	DrdaCpAccrdbrm:    "ACCRDBRM",
	DrdaCpQrynoprm:    "QRYNOPRM",
	DrdaCpRdbnacrm:    "RDBNACRM",
	DrdaCpOpnqryrm:    "OPNQRYRM",
	DrdaCpPkgbnarm:    "PKGBNARM",
	DrdaCpRdbaccrm:    "RDBACCRM",
	DrdaCpBgnbndrm:    "BGNBNDRM",
	DrdaCpPkgbparm:    "PKGBPARM",
	DrdaCpDscinvrm:    "DSCINVRM",
	DrdaCpEndqryrm:    "ENDQRYRM",
	DrdaCpEnduowrm:    "ENDUOWRM",
	DrdaCpAbnuowrm:    "ABNUOWRM",
	DrdaCpDtamchrm:    "DTAMCHRM",
	DrdaCpQrypoprm:    "QRYPOPRM",
	DrdaCpRdbnfnrm:    "RDBNFNRM",
	DrdaCpOpnqflrm:    "OPNQFLRM",
	DrdaCpSqlerrrm:    "SQLERRRM",
	DrdaCpRdbupdrm:    "RDBUPDRM",
	DrdaCpRslsetrm:    "RSLSETRM",
	DrdaCpRdbaflrm:    "RDBAFLRM",
	DrdaCpCmdvltrm:    "CMDVLTRM",
	DrdaCpCmmrqsrm:    "CMMRQSRM",
	DrdaCpRdbathrm:    "RDBATHRM",
	DrdaCpSqlam:       "SQLAM",
	DrdaCpSqlcard:     "SQLCARD",
	DrdaCpSqlcinrd:    "SQLCINRD",
	DrdaCpSqlrslrd:    "SQLRSLRD",
	DrdaCpRdb:         "RDB",
	DrdaCpFrcfixrow:   "FRCFIXROW",
	DrdaCpSqldard:     "SQLDARD",
	DrdaCpSqldta:      "SQLDTA",
	DrdaCpSqldtard:    "SQLDTARD",
	DrdaCpSqlstt:      "SQLSTT",
	DrdaCpOutovr:      "OUTOVR",
	DrdaCpLmtblkprc:   "LMTBLKPRC",
	DrdaCpFixrowprc:   "FIXROWPRC",
	DrdaCpSqlsttvrb:   "SQLSTTVRB",
	DrdaCpQrydsc:      "QRYDSC",
	DrdaCpQrydta:      "QRYDTA",
	DrdaCpSqlattr:     "SQLATTR",
	DrdaCpPbsd:        "PBSD",

	// Isolation level as a byte (product-specific)
	DrdaCpPbsdIso: "PBSDISO",

	// Current schema as UTF8 String (product-specific)
	DrdaCpPbsdSchema: "PBSD_SCHEMA",
}

var DSS_Abbrev = map[uint16]string{
	DrdaDssfmtRqsdss:   "RQSDSS",
	DrdaDssfmtRpydss:   "RPYDSS",
	DrdaDssfmtObjdss:   "OBJDSS",
	DrdaDssfmtCmndss:   "CMNDSS",
	DrdaDssfmtNorpydss: "NORPYDSS",
	0:                  "NULL",
}
