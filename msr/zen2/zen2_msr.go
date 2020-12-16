package zen2

const (
	EVENT_ACTUAL_CPU_CLOCK = 0x01
	UMASK_ACTUAL_CPU_CLOCK = 0x00

	EVENT_APERF = 0x01
	UMASK_APERF = 0x00

	EVENT_MAX_CPU_CLOCK = 0x02
	UMASK_MAX_CPU_CLOCK = 0x00

	EVENT_MPERF = 0x02
	UMASK_MPERF = 0x00

	EVENT_MERGE = 0xFFF
	UMASK_MERGE = 0x00

	EVENT_RETIRED_SSE_AVX_FLOPS              = 0x03
	UMASK_RETIRED_SSE_AVX_FLOPS_ADD_SUB      = 0x01
	UMASK_RETIRED_SSE_AVX_FLOPS_MULT         = 0x02
	UMASK_RETIRED_SSE_AVX_FLOPS_DIV          = 0x04
	UMASK_RETIRED_SSE_AVX_FLOPS_ADD_MULT_DIV = 0x07
	UMASK_RETIRED_SSE_AVX_FLOPS_FMA          = 0x08
	UMASK_RETIRED_SSE_AVX_FLOPS_ALL          = 0x0F

	EVENT_RETIRED_SERIALIZING_OPS             = 0x05
	UMASK_RETIRED_SERIALIZING_OPS_SSE_BOTTOM  = 0x00
	UMASK_RETIRED_SERIALIZING_OPS_SSE_CONTROL = 0x01
	UMASK_RETIRED_SERIALIZING_OPS_X87_BOTTOM  = 0x02
	UMASK_RETIRED_SERIALIZING_OPS_X87_CONTROL = 0x04

	EVENT_FP_DISPATCH_FAULTS           = 0x0E
	UMASK_FP_DISPATCH_FAULTS_X86_FILL  = 0x01
	UMASK_FP_DISPATCH_FAULTS_XMM_FILL  = 0x02
	UMASK_FP_DISPATCH_FAULTS_YMM_FILL  = 0x04
	UMASK_FP_DISPATCH_FAULTS_YMM_SPILL = 0x08

	EVENT_RETIRED_LOCK_INSTR = 0x25
	UMASK_RETIRED_LOCK_INSTR = 0x00

	EVENT_RETIRED_CLFLUSH = 0x26
	UMASK_RETIRED_CLFLUSH = 0x00

	EVENT_RETIRED_CPUID = 0x27
	UMASK_RETIRED_CPUID = 0x00

	EVENT_LS_DISPATCH                = 0x29
	UMASK_LS_DISPATCH_LOADS          = 0x01
	UMASK_LS_DISPATCH_STORES         = 0x02
	UMASK_LS_DISPATCH_LOAD_OP_STORES = 0x04

	EVENT_SMIS_RECEIVED = 0x2B
	UMASK_SMIS_RECEIVED = 0x00

	EVENT_INTERRUPTS_TAKEN = 0x2C
	UMASK_INTERRUPTS_TAKEN = 0x00

	EVENT_ST_TO_LD_FWD = 0x35
	UMASK_ST_TO_LD_FWD = 0x00

	EVENT_ST_COMMIT_CANCELS = 0x37
	UMASK_ST_COMMIT_CANCELS = 0x01

	EVENT_DATA_CACHE_ACCESSES = 0x40
	UMASK_DATA_CACHE_ACCESSES = 0x00

	EVENT_LS_MAB_ALLOC        = 0x41
	UMASK_LS_MAB_ALLOC_LOADS  = 0x01
	UMASK_LS_MAB_ALLOC_STORES = 0x02
	UMASK_LS_MAB_ALLOC_HW_PF  = 0x08

	EVENT_DATA_CACHE_REFILLS              = 0x43
	UMASK_DATA_CACHE_REFILLS_LOCAL_L2     = 0x01
	UMASK_DATA_CACHE_REFILLS_LOCAL_CACHE  = 0x02
	UMASK_DATA_CACHE_REFILLS_LOCAL_DRAM   = 0x08
	UMASK_DATA_CACHE_REFILLS_LOCAL_ALL    = 0x0B
	UMASK_DATA_CACHE_REFILLS_REMOTE_CACHE = 0x10
	UMASK_DATA_CACHE_REFILLS_REMOTE_DRAM  = 0x40
	UMASK_DATA_CACHE_REFILLS_REMOTE_ALL   = 0x50
	UMASK_DATA_CACHE_REFILLS_ALL          = 0x5B

	EVENT_L1_DTLB_MISS             = 0x45
	UMASK_L1_DTLB_MISS_4K_L2_HIT   = 0x01
	UMASK_L1_DTLB_MISS_32K_L2_HIT  = 0x02
	UMASK_L1_DTLB_MISS_2M_L2_HIT   = 0x04
	UMASK_L1_DTLB_MISS_1G_L2_HIT   = 0x08
	UMASK_L1_DTLB_MISS_ANY_L2_HIT  = 0x0F
	UMASK_L1_DTLB_MISS_4K_L2_MISS  = 0x10
	UMASK_L1_DTLB_MISS_32K_L2_MISS = 0x20
	UMASK_L1_DTLB_MISS_2M_L2_MISS  = 0x40
	UMASK_L1_DTLB_MISS_1G_L2_MISS  = 0x80
	UMASK_L1_DTLB_MISS_ANY_L2_MISS = 0xF0

	EVENT_TABLEWALKER_ALLOC        = 0x46
	UMASK_TABLEWALKER_ALLOC_DSIDE0 = 0x01
	UMASK_TABLEWALKER_ALLOC_DSIDE1 = 0x02
	UMASK_TABLEWALKER_ALLOC_ISIDE0 = 0x04
	UMASK_TABLEWALKER_ALLOC_ISIDE1 = 0x08

	EVENT_MISALIGNED_LOADS = 0x47
	UMASK_MISALIGNED_LOADS = 0x00

	EVENT_PREF_INSTR_DISPATCHED = 0x4B
	UMASK_PREF_INSTR_DISPATCHED = 0x00

	EVENT_INEFFECTIVE_SW_PREF                = 0x52
	UMASK_INEFFECTIVE_SW_PREF_DATA_CACHE_HIT = 0x01
	UMASK_INEFFECTIVE_SW_PREF_MAB_MATCH      = 0x02

	EVENT_SWPREF_DATA_CACHE_FILLS              = 0x59
	UMASK_SWPREF_DATA_CACHE_FILLS_LOCAL_L2     = 0x01
	UMASK_SWPREF_DATA_CACHE_FILLS_LOCAL_CACHE  = 0x02
	UMASK_SWPREF_DATA_CACHE_FILLS_LOCAL_DRAM   = 0x08
	UMASK_SWPREF_DATA_CACHE_FILLS_LOCAL_ALL    = 0x0B
	UMASK_SWPREF_DATA_CACHE_FILLS_REMOTE_CACHE = 0x10
	UMASK_SWPREF_DATA_CACHE_FILLS_REMOTE_DRAM  = 0x40
	UMASK_SWPREF_DATA_CACHE_FILLS_REMOTE_ALL   = 0x50
	UMASK_SWPREF_DATA_CACHE_FILLS_ALL          = 0x5B

	EVENT_HWPREF_DATA_CACHE_FILLS              = 0x5A
	UMASK_HWPREF_DATA_CACHE_FILLS_LOCAL_L2     = 0x01
	UMASK_HWPREF_DATA_CACHE_FILLS_LOCAL_CACHE  = 0x02
	UMASK_HWPREF_DATA_CACHE_FILLS_LOCAL_DRAM   = 0x08
	UMASK_HWPREF_DATA_CACHE_FILLS_LOCAL_ALL    = 0x0B
	UMASK_HWPREF_DATA_CACHE_FILLS_REMOTE_CACHE = 0x10
	UMASK_HWPREF_DATA_CACHE_FILLS_REMOTE_DRAM  = 0x40
	UMASK_HWPREF_DATA_CACHE_FILLS_REMOTE_ALL   = 0x50
	UMASK_HWPREF_DATA_CACHE_FILLS_ALL          = 0x5B

	EVENT_CPU_CLOCKS_UNHALTED = 0x76
	UMASK_CPU_CLOCKS_UNHALTED = 0x00

	EVENT_TLB_FLUSHES = 0x78
	UMASK_TLB_FLUSHES = 0x00

	EVENT_ICACHE_FETCHES = 0x80
	UMASK_ICACHE_FETCHES = 0x00

	EVENT_ICACHE_MISSES = 0x81
	UMASK_ICACHE_MISSES = 0x00

	EVENT_ICACHE_L2_REFILLS = 0x82
	UMASK_ICACHE_L2_REFILLS = 0x00

	EVENT_ICACHE_SYSTEM_REFILLS = 0x83
	UMASK_ICACHE_SYSTEM_REFILLS = 0x00

	EVENT_L1_ITLB_MISS_L2_ITLB_HIT = 0x84
	UMASK_L1_ITLB_MISS_L2_ITLB_HIT = 0x00

	EVENT_L1_ITLB_MISS_L2_ITLB_MISS = 0x85
	UMASK_L1_ITLB_MISS_L2_ITLB_MISS = 0x00

	EVENT_L1_BTB_CORRECTION = 0x8A
	UMASK_L1_BTB_CORRECTION = 0x00

	EVENT_L2_BTB_CORRECTION = 0x8B
	UMASK_L2_BTB_CORRECTION = 0x00

	EVENT_ICACHE_LINES_INVALIDATED          = 0x8C
	UMASK_ICACHE_LINES_INVALIDATED_FILL     = 0x01
	UMASK_ICACHE_LINES_INVALIDATED_L2_PROBE = 0x02

	EVENT_DEC_OVERRIDE_BTB = 0x91
	UMASK_DEC_OVERRIDE_BTB = 0x00

	EVENT_UOPS_DISP              = 0xAA
	UMASK_UOPS_DISP_FROM_DEC     = 0x01
	UMASK_UOPS_DISP_FROM_OPCACHE = 0x02

	EVENT_DYN_TOKENS_DISP_STALL_CYCLES0                 = 0xAF
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES0_ALU_TOKEN_STALL = 0x08

	EVENT_DYN_TOKENS_DISP_STALL_CYCLES1                           = 0xAE
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_INT_REG_FILE_STALL        = 0x01
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_LD_QUEUE_STALL            = 0x02
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_ST_QUEUE_STALL            = 0x04
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_INT_SCHED_MISC_STALL      = 0x08
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_TAKEN_BRANCH_BUFFER_STALL = 0x10
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_FP_REG_FILE_STALL         = 0x20
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_FP_SCHED_STALL            = 0x40
	UMASK_DYN_TOKENS_DISP_STALL_CYCLES1_FP_MISC_UNAVAIL           = 0x80

	EVENT_RETIRED_INSTRUCTIONS = 0xC0
	UMASK_RETIRED_INSTRUCTIONS = 0x00

	EVENT_RETIRED_UOPS = 0xC1
	UMASK_RETIRED_UOPS = 0x00

	EVENT_RETIRED_BRANCH_INSTR = 0xC2
	UMASK_RETIRED_BRANCH_INSTR = 0x00

	EVENT_RETIRED_MISP_BRANCH_INSTR = 0xC3
	UMASK_RETIRED_MISP_BRANCH_INSTR = 0x00

	EVENT_RETIRED_TAKEN_BRANCH_INSTR = 0xC4
	UMASK_RETIRED_TAKEN_BRANCH_INSTR = 0x00

	EVENT_RETIRED_TAKEN_MISP_BRANCH_INSTR = 0xC5
	UMASK_RETIRED_TAKEN_MISP_BRANCH_INSTR = 0x00

	EVENT_RETIRED_FAR_CONTROL_TRANSFERS = 0xC6
	UMASK_RETIRED_FAR_CONTROL_TRANSFERS = 0x00

	EVENT_RETIRED_NEAR_RETURNS = 0xC8
	UMASK_RETIRED_NEAR_RETURNS = 0x00

	EVENT_RETIRED_NEAR_RETURNS_MISP = 0xC9
	UMASK_RETIRED_NEAR_RETURNS_MISP = 0x00

	EVENT_RETIRED_INDIRECT_BRANCHES_MISP = 0xCA
	UMASK_RETIRED_INDIRECT_BRANCHES_MISP = 0x00

	EVENT_RETIRED_MMX_FP_INSTR     = 0xCB
	UMASK_RETIRED_MMX_FP_INSTR_X87 = 0x01
	UMASK_RETIRED_MMX_FP_INSTR_MMX = 0x02
	UMASK_RETIRED_MMX_FP_INSTR_SSE = 0x04
	UMASK_RETIRED_MMX_FP_INSTR_ALL = 0x07

	EVENT_RETIRED_COND_BRANCH_INSTR = 0xD1
	UMASK_RETIRED_COND_BRANCH_INSTR = 0x00

	EVENT_DIV_BUSY_CYCLES = 0xD3
	UMASK_DIV_BUSY_CYCLES = 0x00

	EVENT_DIV_OP_COUNT = 0xD4
	UMASK_DIV_OP_COUNT = 0x00

	EVENT_TAGGED_IBS_OPS                = 0x1CF
	UMASK_TAGGED_IBS_OPS_COUNT          = 0x01
	UMASK_TAGGED_IBS_OPS_COUNT_RETIRED  = 0x02
	UMASK_TAGGED_IBS_OPS_COUNT_ROLLOVER = 0x04

	EVENT_RETIRED_FUSED_BRANCH_INSTR = 0x1D0
	UMASK_RETIRED_FUSED_BRANCH_INSTR = 0x00

	EVENT_REQUESTS_TO_L2_GRP1                   = 0x60
	UMASK_REQUESTS_TO_L2_GRP1_GRP2              = 0x01
	UMASK_REQUESTS_TO_L2_GRP1_L2_HW_PREF        = 0x02
	UMASK_REQUESTS_TO_L2_GRP1_PREF_L2           = 0x04
	UMASK_REQUESTS_TO_L2_GRP1_CHANGE_TO_X       = 0x08
	UMASK_REQUESTS_TO_L2_GRP1_CACHEABLE_IC_READ = 0x10
	UMASK_REQUESTS_TO_L2_GRP1_LS_RD_BLOCK_C_S   = 0x20
	UMASK_REQUESTS_TO_L2_GRP1_RD_BLOCK_X        = 0x40
	UMASK_REQUESTS_TO_L2_GRP1_RD_BLOCK_L        = 0x80
	UMASK_REQUESTS_TO_L2_GRP1_DATA_CACHE_MISSES = 0xC8
	UMASK_REQUESTS_TO_L2_GRP1_L1_CACHES_MISS    = 0x88
	UMASK_REQUESTS_TO_L2_GRP1_ALL_NO_PF         = 0xF9
	UMASK_REQUESTS_TO_L2_GRP1_ALL               = 0xFF

	EVENT_REQUESTS_TO_L2_GRP2                  = 0x61
	UMASK_REQUESTS_TO_L2_GRP2_BUS_LOCK_RESP    = 0x01
	UMASK_REQUESTS_TO_L2_GRP2_BUS_LOCK_ORIG    = 0x02
	UMASK_REQUESTS_TO_L2_GRP2_SMC_INVAL        = 0x04
	UMASK_REQUESTS_TO_L2_GRP2_IC_READ_SIZED_NC = 0x08
	UMASK_REQUESTS_TO_L2_GRP2_IC_READ_SIZED    = 0x10
	UMASK_REQUESTS_TO_L2_GRP2_IC_READ          = 0x18
	UMASK_REQUESTS_TO_L2_GRP2_LS_READ_SIZED_NC = 0x20
	UMASK_REQUESTS_TO_L2_GRP2_LS_READ_SIZED    = 0x40
	UMASK_REQUESTS_TO_L2_GRP2_LS_READ          = 0x60
	UMASK_REQUESTS_TO_L2_GRP2_GRP1             = 0x80

	EVENT_CORE_TO_L2_CACHE_REQUESTS                     = 0x64
	UMASK_CORE_TO_L2_CACHE_REQUESTS_IC_FILL_MISS        = 0x01
	UMASK_CORE_TO_L2_CACHE_REQUESTS_IC_FILL_HIT_S       = 0x02
	UMASK_CORE_TO_L2_CACHE_REQUESTS_IC_FILL_HIT_X       = 0x04
	UMASK_CORE_TO_L2_CACHE_REQUESTS_LD_READ_BLK_C       = 0x08
	UMASK_CORE_TO_L2_CACHE_REQUESTS_LD_READ_BLK_X       = 0x10
	UMASK_CORE_TO_L2_CACHE_REQUESTS_LD_READ_BLK_L_HIT_S = 0x20
	UMASK_CORE_TO_L2_CACHE_REQUESTS_LD_READ_BLK_L_HIT_X = 0x40
	UMASK_CORE_TO_L2_CACHE_REQUESTS_LD_READ_BLK_CS      = 0x80

	EVENT_L2_PF_HIT_IN_L2 = 0x70
	UMASK_L2_PF_HIT_IN_L2 = 0x1F

	EVENT_L2_PF_HIT_IN_L3 = 0x71
	UMASK_L2_PF_HIT_IN_L3 = 0x1F

	EVENT_L2_PF_MISS_IN_L3 = 0x72
	UMASK_L2_PF_MISS_IN_L3 = 0x1F

	EVENT_L3_ACCESS = 0x01
	UMASK_L3_ACCESS = 0x80

	EVENT_L3_MISS = 0x06
	UMASK_L3_MISS = 0x01

	EVENT_L3_CACHE_REQ = 0x04
	UMASK_L3_CACHE_REQ = 0xFF

	EVENT_L3_MISS_LAT = 0x90
	UMASK_L3_MISS_LAT = 0x00

	EVENT_L3_MISS_REQ = 0x9A
	UMASK_L3_MISS_REQ = 0x1F

	EVENT_RAPL_CORE_ENERGY = 0x01
	UMASK_RAPL_CORE_ENERGY = 0x00

	EVENT_RAPL_PKG_ENERGY = 0x02
	UMASK_RAPL_PKG_ENERGY = 0x00

	EVENT_DATA_FROM_LOCAL_DRAM_CHANNEL = 0x07
	UMASK_DATA_FROM_LOCAL_DRAM_CHANNEL = 0x38

	EVENT_DATA_TO_LOCAL_DRAM_CHANNEL = 0x47
	UMASK_DATA_TO_LOCAL_DRAM_CHANNEL = 0x38

	EVENT_DATA_OUT_TO_REMOTE_0 = 0x187
	UMASK_DATA_OUT_TO_REMOTE_0 = 0x02

	EVENT_DATA_OUT_TO_REMOTE_1 = 0x1C7
	UMASK_DATA_OUT_TO_REMOTE_1 = 0x02

	EVENT_DATA_OUT_TO_REMOTE_2 = 0x207
	UMASK_DATA_OUT_TO_REMOTE_2 = 0x02

	EVENT_DATA_OUT_TO_REMOTE_3 = 0x287
	UMASK_DATA_OUT_TO_REMOTE_3 = 0x02

	EVENT_DATA_OUT_TO_REMOTE_4 = 0x247
	UMASK_DATA_OUT_TO_REMOTE_4 = 0x02

	EVENT_DATA_OUT_TO_REMOTE_5 = 0x2C7
	UMASK_DATA_OUT_TO_REMOTE_5 = 0x02
)
