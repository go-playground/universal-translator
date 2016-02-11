package locales

import (
	"errors"

	"github.com/go-playground/universal-translator"
	"github.com/go-playground/universal-translator/resources/locales/af"
	"github.com/go-playground/universal-translator/resources/locales/af_NA"
	"github.com/go-playground/universal-translator/resources/locales/agq"
	"github.com/go-playground/universal-translator/resources/locales/ak"
	"github.com/go-playground/universal-translator/resources/locales/am"
	"github.com/go-playground/universal-translator/resources/locales/ar"
	"github.com/go-playground/universal-translator/resources/locales/ar_AE"
	"github.com/go-playground/universal-translator/resources/locales/ar_DJ"
	"github.com/go-playground/universal-translator/resources/locales/ar_DZ"
	"github.com/go-playground/universal-translator/resources/locales/ar_EH"
	"github.com/go-playground/universal-translator/resources/locales/ar_ER"
	"github.com/go-playground/universal-translator/resources/locales/ar_LB"
	"github.com/go-playground/universal-translator/resources/locales/ar_LY"
	"github.com/go-playground/universal-translator/resources/locales/ar_MA"
	"github.com/go-playground/universal-translator/resources/locales/ar_MR"
	"github.com/go-playground/universal-translator/resources/locales/ar_SO"
	"github.com/go-playground/universal-translator/resources/locales/ar_SS"
	"github.com/go-playground/universal-translator/resources/locales/ar_TN"
	"github.com/go-playground/universal-translator/resources/locales/as"
	"github.com/go-playground/universal-translator/resources/locales/asa"
	"github.com/go-playground/universal-translator/resources/locales/ast"
	"github.com/go-playground/universal-translator/resources/locales/az"
	"github.com/go-playground/universal-translator/resources/locales/az_Cyrl"
	"github.com/go-playground/universal-translator/resources/locales/bas"
	"github.com/go-playground/universal-translator/resources/locales/be"
	"github.com/go-playground/universal-translator/resources/locales/bem"
	"github.com/go-playground/universal-translator/resources/locales/bez"
	"github.com/go-playground/universal-translator/resources/locales/bg"
	"github.com/go-playground/universal-translator/resources/locales/bm"
	"github.com/go-playground/universal-translator/resources/locales/bn"
	"github.com/go-playground/universal-translator/resources/locales/bo"
	"github.com/go-playground/universal-translator/resources/locales/bo_IN"
	"github.com/go-playground/universal-translator/resources/locales/br"
	"github.com/go-playground/universal-translator/resources/locales/brx"
	"github.com/go-playground/universal-translator/resources/locales/bs"
	"github.com/go-playground/universal-translator/resources/locales/bs_Cyrl"
	"github.com/go-playground/universal-translator/resources/locales/ca"
	"github.com/go-playground/universal-translator/resources/locales/ca_ES_VALENCIA"
	"github.com/go-playground/universal-translator/resources/locales/ca_FR"
	"github.com/go-playground/universal-translator/resources/locales/ce"
	"github.com/go-playground/universal-translator/resources/locales/cgg"
	"github.com/go-playground/universal-translator/resources/locales/chr"
	"github.com/go-playground/universal-translator/resources/locales/ckb"
	"github.com/go-playground/universal-translator/resources/locales/cs"
	"github.com/go-playground/universal-translator/resources/locales/cu"
	"github.com/go-playground/universal-translator/resources/locales/cy"
	"github.com/go-playground/universal-translator/resources/locales/da"
	"github.com/go-playground/universal-translator/resources/locales/dav"
	"github.com/go-playground/universal-translator/resources/locales/de"
	"github.com/go-playground/universal-translator/resources/locales/de_AT"
	"github.com/go-playground/universal-translator/resources/locales/de_CH"
	"github.com/go-playground/universal-translator/resources/locales/de_LI"
	"github.com/go-playground/universal-translator/resources/locales/de_LU"
	"github.com/go-playground/universal-translator/resources/locales/dje"
	"github.com/go-playground/universal-translator/resources/locales/dsb"
	"github.com/go-playground/universal-translator/resources/locales/dua"
	"github.com/go-playground/universal-translator/resources/locales/dyo"
	"github.com/go-playground/universal-translator/resources/locales/dz"
	"github.com/go-playground/universal-translator/resources/locales/ebu"
	"github.com/go-playground/universal-translator/resources/locales/ee"
	"github.com/go-playground/universal-translator/resources/locales/el"
	"github.com/go-playground/universal-translator/resources/locales/en"
	"github.com/go-playground/universal-translator/resources/locales/en_001"
	"github.com/go-playground/universal-translator/resources/locales/en_150"
	"github.com/go-playground/universal-translator/resources/locales/en_AG"
	"github.com/go-playground/universal-translator/resources/locales/en_AI"
	"github.com/go-playground/universal-translator/resources/locales/en_AT"
	"github.com/go-playground/universal-translator/resources/locales/en_AU"
	"github.com/go-playground/universal-translator/resources/locales/en_BB"
	"github.com/go-playground/universal-translator/resources/locales/en_BE"
	"github.com/go-playground/universal-translator/resources/locales/en_BI"
	"github.com/go-playground/universal-translator/resources/locales/en_BM"
	"github.com/go-playground/universal-translator/resources/locales/en_BS"
	"github.com/go-playground/universal-translator/resources/locales/en_BW"
	"github.com/go-playground/universal-translator/resources/locales/en_BZ"
	"github.com/go-playground/universal-translator/resources/locales/en_CA"
	"github.com/go-playground/universal-translator/resources/locales/en_CC"
	"github.com/go-playground/universal-translator/resources/locales/en_CH"
	"github.com/go-playground/universal-translator/resources/locales/en_CK"
	"github.com/go-playground/universal-translator/resources/locales/en_CX"
	"github.com/go-playground/universal-translator/resources/locales/en_DE"
	"github.com/go-playground/universal-translator/resources/locales/en_DK"
	"github.com/go-playground/universal-translator/resources/locales/en_DM"
	"github.com/go-playground/universal-translator/resources/locales/en_ER"
	"github.com/go-playground/universal-translator/resources/locales/en_FI"
	"github.com/go-playground/universal-translator/resources/locales/en_FJ"
	"github.com/go-playground/universal-translator/resources/locales/en_FK"
	"github.com/go-playground/universal-translator/resources/locales/en_GB"
	"github.com/go-playground/universal-translator/resources/locales/en_GD"
	"github.com/go-playground/universal-translator/resources/locales/en_GG"
	"github.com/go-playground/universal-translator/resources/locales/en_GH"
	"github.com/go-playground/universal-translator/resources/locales/en_GI"
	"github.com/go-playground/universal-translator/resources/locales/en_GM"
	"github.com/go-playground/universal-translator/resources/locales/en_GY"
	"github.com/go-playground/universal-translator/resources/locales/en_IM"
	"github.com/go-playground/universal-translator/resources/locales/en_IN"
	"github.com/go-playground/universal-translator/resources/locales/en_JE"
	"github.com/go-playground/universal-translator/resources/locales/en_JM"
	"github.com/go-playground/universal-translator/resources/locales/en_KE"
	"github.com/go-playground/universal-translator/resources/locales/en_KI"
	"github.com/go-playground/universal-translator/resources/locales/en_KN"
	"github.com/go-playground/universal-translator/resources/locales/en_KY"
	"github.com/go-playground/universal-translator/resources/locales/en_LC"
	"github.com/go-playground/universal-translator/resources/locales/en_LR"
	"github.com/go-playground/universal-translator/resources/locales/en_LS"
	"github.com/go-playground/universal-translator/resources/locales/en_MG"
	"github.com/go-playground/universal-translator/resources/locales/en_MO"
	"github.com/go-playground/universal-translator/resources/locales/en_MS"
	"github.com/go-playground/universal-translator/resources/locales/en_MT"
	"github.com/go-playground/universal-translator/resources/locales/en_MU"
	"github.com/go-playground/universal-translator/resources/locales/en_MW"
	"github.com/go-playground/universal-translator/resources/locales/en_MY"
	"github.com/go-playground/universal-translator/resources/locales/en_NA"
	"github.com/go-playground/universal-translator/resources/locales/en_NF"
	"github.com/go-playground/universal-translator/resources/locales/en_NG"
	"github.com/go-playground/universal-translator/resources/locales/en_NL"
	"github.com/go-playground/universal-translator/resources/locales/en_NR"
	"github.com/go-playground/universal-translator/resources/locales/en_NU"
	"github.com/go-playground/universal-translator/resources/locales/en_NZ"
	"github.com/go-playground/universal-translator/resources/locales/en_PG"
	"github.com/go-playground/universal-translator/resources/locales/en_PH"
	"github.com/go-playground/universal-translator/resources/locales/en_PK"
	"github.com/go-playground/universal-translator/resources/locales/en_PN"
	"github.com/go-playground/universal-translator/resources/locales/en_RW"
	"github.com/go-playground/universal-translator/resources/locales/en_SB"
	"github.com/go-playground/universal-translator/resources/locales/en_SC"
	"github.com/go-playground/universal-translator/resources/locales/en_SE"
	"github.com/go-playground/universal-translator/resources/locales/en_SG"
	"github.com/go-playground/universal-translator/resources/locales/en_SH"
	"github.com/go-playground/universal-translator/resources/locales/en_SI"
	"github.com/go-playground/universal-translator/resources/locales/en_SL"
	"github.com/go-playground/universal-translator/resources/locales/en_SS"
	"github.com/go-playground/universal-translator/resources/locales/en_SX"
	"github.com/go-playground/universal-translator/resources/locales/en_SZ"
	"github.com/go-playground/universal-translator/resources/locales/en_TK"
	"github.com/go-playground/universal-translator/resources/locales/en_TO"
	"github.com/go-playground/universal-translator/resources/locales/en_TT"
	"github.com/go-playground/universal-translator/resources/locales/en_TV"
	"github.com/go-playground/universal-translator/resources/locales/en_TZ"
	"github.com/go-playground/universal-translator/resources/locales/en_UG"
	"github.com/go-playground/universal-translator/resources/locales/en_US_POSIX"
	"github.com/go-playground/universal-translator/resources/locales/en_VC"
	"github.com/go-playground/universal-translator/resources/locales/en_VU"
	"github.com/go-playground/universal-translator/resources/locales/en_WS"
	"github.com/go-playground/universal-translator/resources/locales/en_ZA"
	"github.com/go-playground/universal-translator/resources/locales/en_ZM"
	"github.com/go-playground/universal-translator/resources/locales/eo"
	"github.com/go-playground/universal-translator/resources/locales/es"
	"github.com/go-playground/universal-translator/resources/locales/es_419"
	"github.com/go-playground/universal-translator/resources/locales/es_AR"
	"github.com/go-playground/universal-translator/resources/locales/es_BO"
	"github.com/go-playground/universal-translator/resources/locales/es_CL"
	"github.com/go-playground/universal-translator/resources/locales/es_CO"
	"github.com/go-playground/universal-translator/resources/locales/es_CR"
	"github.com/go-playground/universal-translator/resources/locales/es_CU"
	"github.com/go-playground/universal-translator/resources/locales/es_DO"
	"github.com/go-playground/universal-translator/resources/locales/es_EC"
	"github.com/go-playground/universal-translator/resources/locales/es_GQ"
	"github.com/go-playground/universal-translator/resources/locales/es_GT"
	"github.com/go-playground/universal-translator/resources/locales/es_HN"
	"github.com/go-playground/universal-translator/resources/locales/es_MX"
	"github.com/go-playground/universal-translator/resources/locales/es_NI"
	"github.com/go-playground/universal-translator/resources/locales/es_PA"
	"github.com/go-playground/universal-translator/resources/locales/es_PE"
	"github.com/go-playground/universal-translator/resources/locales/es_PH"
	"github.com/go-playground/universal-translator/resources/locales/es_PR"
	"github.com/go-playground/universal-translator/resources/locales/es_PY"
	"github.com/go-playground/universal-translator/resources/locales/es_SV"
	"github.com/go-playground/universal-translator/resources/locales/es_US"
	"github.com/go-playground/universal-translator/resources/locales/es_UY"
	"github.com/go-playground/universal-translator/resources/locales/es_VE"
	"github.com/go-playground/universal-translator/resources/locales/et"
	"github.com/go-playground/universal-translator/resources/locales/eu"
	"github.com/go-playground/universal-translator/resources/locales/ewo"
	"github.com/go-playground/universal-translator/resources/locales/fa"
	"github.com/go-playground/universal-translator/resources/locales/fa_AF"
	"github.com/go-playground/universal-translator/resources/locales/ff"
	"github.com/go-playground/universal-translator/resources/locales/ff_GN"
	"github.com/go-playground/universal-translator/resources/locales/ff_MR"
	"github.com/go-playground/universal-translator/resources/locales/fi"
	"github.com/go-playground/universal-translator/resources/locales/fil"
	"github.com/go-playground/universal-translator/resources/locales/fo"
	"github.com/go-playground/universal-translator/resources/locales/fo_DK"
	"github.com/go-playground/universal-translator/resources/locales/fr"
	"github.com/go-playground/universal-translator/resources/locales/fr_BE"
	"github.com/go-playground/universal-translator/resources/locales/fr_BI"
	"github.com/go-playground/universal-translator/resources/locales/fr_CA"
	"github.com/go-playground/universal-translator/resources/locales/fr_CD"
	"github.com/go-playground/universal-translator/resources/locales/fr_CH"
	"github.com/go-playground/universal-translator/resources/locales/fr_DJ"
	"github.com/go-playground/universal-translator/resources/locales/fr_DZ"
	"github.com/go-playground/universal-translator/resources/locales/fr_GN"
	"github.com/go-playground/universal-translator/resources/locales/fr_HT"
	"github.com/go-playground/universal-translator/resources/locales/fr_KM"
	"github.com/go-playground/universal-translator/resources/locales/fr_LU"
	"github.com/go-playground/universal-translator/resources/locales/fr_MA"
	"github.com/go-playground/universal-translator/resources/locales/fr_MG"
	"github.com/go-playground/universal-translator/resources/locales/fr_MR"
	"github.com/go-playground/universal-translator/resources/locales/fr_MU"
	"github.com/go-playground/universal-translator/resources/locales/fr_RW"
	"github.com/go-playground/universal-translator/resources/locales/fr_SC"
	"github.com/go-playground/universal-translator/resources/locales/fr_SY"
	"github.com/go-playground/universal-translator/resources/locales/fr_TN"
	"github.com/go-playground/universal-translator/resources/locales/fr_VU"
	"github.com/go-playground/universal-translator/resources/locales/fur"
	"github.com/go-playground/universal-translator/resources/locales/fy"
	"github.com/go-playground/universal-translator/resources/locales/ga"
	"github.com/go-playground/universal-translator/resources/locales/gd"
	"github.com/go-playground/universal-translator/resources/locales/gl"
	"github.com/go-playground/universal-translator/resources/locales/gsw"
	"github.com/go-playground/universal-translator/resources/locales/gu"
	"github.com/go-playground/universal-translator/resources/locales/guz"
	"github.com/go-playground/universal-translator/resources/locales/gv"
	"github.com/go-playground/universal-translator/resources/locales/ha"
	"github.com/go-playground/universal-translator/resources/locales/ha_GH"
	"github.com/go-playground/universal-translator/resources/locales/haw"
	"github.com/go-playground/universal-translator/resources/locales/he"
	"github.com/go-playground/universal-translator/resources/locales/hi"
	"github.com/go-playground/universal-translator/resources/locales/hr"
	"github.com/go-playground/universal-translator/resources/locales/hr_BA"
	"github.com/go-playground/universal-translator/resources/locales/hsb"
	"github.com/go-playground/universal-translator/resources/locales/hu"
	"github.com/go-playground/universal-translator/resources/locales/hy"
	"github.com/go-playground/universal-translator/resources/locales/id"
	"github.com/go-playground/universal-translator/resources/locales/ig"
	"github.com/go-playground/universal-translator/resources/locales/ii"
	"github.com/go-playground/universal-translator/resources/locales/is"
	"github.com/go-playground/universal-translator/resources/locales/it"
	"github.com/go-playground/universal-translator/resources/locales/it_CH"
	"github.com/go-playground/universal-translator/resources/locales/ja"
	"github.com/go-playground/universal-translator/resources/locales/jgo"
	"github.com/go-playground/universal-translator/resources/locales/jmc"
	"github.com/go-playground/universal-translator/resources/locales/ka"
	"github.com/go-playground/universal-translator/resources/locales/kab"
	"github.com/go-playground/universal-translator/resources/locales/kam"
	"github.com/go-playground/universal-translator/resources/locales/kde"
	"github.com/go-playground/universal-translator/resources/locales/kea"
	"github.com/go-playground/universal-translator/resources/locales/khq"
	"github.com/go-playground/universal-translator/resources/locales/ki"
	"github.com/go-playground/universal-translator/resources/locales/kk"
	"github.com/go-playground/universal-translator/resources/locales/kkj"
	"github.com/go-playground/universal-translator/resources/locales/kl"
	"github.com/go-playground/universal-translator/resources/locales/kln"
	"github.com/go-playground/universal-translator/resources/locales/km"
	"github.com/go-playground/universal-translator/resources/locales/kn"
	"github.com/go-playground/universal-translator/resources/locales/ko"
	"github.com/go-playground/universal-translator/resources/locales/kok"
	"github.com/go-playground/universal-translator/resources/locales/ks"
	"github.com/go-playground/universal-translator/resources/locales/ksb"
	"github.com/go-playground/universal-translator/resources/locales/ksf"
	"github.com/go-playground/universal-translator/resources/locales/ksh"
	"github.com/go-playground/universal-translator/resources/locales/kw"
	"github.com/go-playground/universal-translator/resources/locales/ky"
	"github.com/go-playground/universal-translator/resources/locales/lag"
	"github.com/go-playground/universal-translator/resources/locales/lb"
	"github.com/go-playground/universal-translator/resources/locales/lg"
	"github.com/go-playground/universal-translator/resources/locales/lkt"
	"github.com/go-playground/universal-translator/resources/locales/ln"
	"github.com/go-playground/universal-translator/resources/locales/ln_AO"
	"github.com/go-playground/universal-translator/resources/locales/lo"
	"github.com/go-playground/universal-translator/resources/locales/lrc"
	"github.com/go-playground/universal-translator/resources/locales/lt"
	"github.com/go-playground/universal-translator/resources/locales/lu"
	"github.com/go-playground/universal-translator/resources/locales/luo"
	"github.com/go-playground/universal-translator/resources/locales/luy"
	"github.com/go-playground/universal-translator/resources/locales/lv"
	"github.com/go-playground/universal-translator/resources/locales/mas"
	"github.com/go-playground/universal-translator/resources/locales/mas_TZ"
	"github.com/go-playground/universal-translator/resources/locales/mer"
	"github.com/go-playground/universal-translator/resources/locales/mfe"
	"github.com/go-playground/universal-translator/resources/locales/mg"
	"github.com/go-playground/universal-translator/resources/locales/mgh"
	"github.com/go-playground/universal-translator/resources/locales/mgo"
	"github.com/go-playground/universal-translator/resources/locales/mk"
	"github.com/go-playground/universal-translator/resources/locales/ml"
	"github.com/go-playground/universal-translator/resources/locales/mn"
	"github.com/go-playground/universal-translator/resources/locales/mr"
	"github.com/go-playground/universal-translator/resources/locales/ms"
	"github.com/go-playground/universal-translator/resources/locales/ms_BN"
	"github.com/go-playground/universal-translator/resources/locales/ms_SG"
	"github.com/go-playground/universal-translator/resources/locales/mt"
	"github.com/go-playground/universal-translator/resources/locales/mua"
	"github.com/go-playground/universal-translator/resources/locales/my"
	"github.com/go-playground/universal-translator/resources/locales/mzn"
	"github.com/go-playground/universal-translator/resources/locales/naq"
	"github.com/go-playground/universal-translator/resources/locales/nb"
	"github.com/go-playground/universal-translator/resources/locales/nd"
	"github.com/go-playground/universal-translator/resources/locales/ne"
	"github.com/go-playground/universal-translator/resources/locales/nl"
	"github.com/go-playground/universal-translator/resources/locales/nl_AW"
	"github.com/go-playground/universal-translator/resources/locales/nl_BE"
	"github.com/go-playground/universal-translator/resources/locales/nl_BQ"
	"github.com/go-playground/universal-translator/resources/locales/nl_CW"
	"github.com/go-playground/universal-translator/resources/locales/nl_SR"
	"github.com/go-playground/universal-translator/resources/locales/nl_SX"
	"github.com/go-playground/universal-translator/resources/locales/nmg"
	"github.com/go-playground/universal-translator/resources/locales/nn"
	"github.com/go-playground/universal-translator/resources/locales/nnh"
	"github.com/go-playground/universal-translator/resources/locales/nus"
	"github.com/go-playground/universal-translator/resources/locales/nyn"
	"github.com/go-playground/universal-translator/resources/locales/om"
	"github.com/go-playground/universal-translator/resources/locales/om_KE"
	"github.com/go-playground/universal-translator/resources/locales/or"
	"github.com/go-playground/universal-translator/resources/locales/os"
	"github.com/go-playground/universal-translator/resources/locales/os_RU"
	"github.com/go-playground/universal-translator/resources/locales/pa"
	"github.com/go-playground/universal-translator/resources/locales/pa_Arab"
	"github.com/go-playground/universal-translator/resources/locales/pl"
	"github.com/go-playground/universal-translator/resources/locales/prg"
	"github.com/go-playground/universal-translator/resources/locales/ps"
	"github.com/go-playground/universal-translator/resources/locales/pt"
	"github.com/go-playground/universal-translator/resources/locales/pt_AO"
	"github.com/go-playground/universal-translator/resources/locales/pt_CV"
	"github.com/go-playground/universal-translator/resources/locales/pt_MO"
	"github.com/go-playground/universal-translator/resources/locales/pt_MZ"
	"github.com/go-playground/universal-translator/resources/locales/pt_PT"
	"github.com/go-playground/universal-translator/resources/locales/pt_ST"
	"github.com/go-playground/universal-translator/resources/locales/qu"
	"github.com/go-playground/universal-translator/resources/locales/qu_BO"
	"github.com/go-playground/universal-translator/resources/locales/qu_EC"
	"github.com/go-playground/universal-translator/resources/locales/rm"
	"github.com/go-playground/universal-translator/resources/locales/rn"
	"github.com/go-playground/universal-translator/resources/locales/ro"
	"github.com/go-playground/universal-translator/resources/locales/ro_MD"
	"github.com/go-playground/universal-translator/resources/locales/rof"
	"github.com/go-playground/universal-translator/resources/locales/root"
	"github.com/go-playground/universal-translator/resources/locales/ru"
	"github.com/go-playground/universal-translator/resources/locales/ru_BY"
	"github.com/go-playground/universal-translator/resources/locales/ru_KG"
	"github.com/go-playground/universal-translator/resources/locales/ru_KZ"
	"github.com/go-playground/universal-translator/resources/locales/ru_MD"
	"github.com/go-playground/universal-translator/resources/locales/rw"
	"github.com/go-playground/universal-translator/resources/locales/rwk"
	"github.com/go-playground/universal-translator/resources/locales/sah"
	"github.com/go-playground/universal-translator/resources/locales/saq"
	"github.com/go-playground/universal-translator/resources/locales/sbp"
	"github.com/go-playground/universal-translator/resources/locales/se"
	"github.com/go-playground/universal-translator/resources/locales/se_SE"
	"github.com/go-playground/universal-translator/resources/locales/seh"
	"github.com/go-playground/universal-translator/resources/locales/ses"
	"github.com/go-playground/universal-translator/resources/locales/sg"
	"github.com/go-playground/universal-translator/resources/locales/shi"
	"github.com/go-playground/universal-translator/resources/locales/shi_Latn"
	"github.com/go-playground/universal-translator/resources/locales/si"
	"github.com/go-playground/universal-translator/resources/locales/sk"
	"github.com/go-playground/universal-translator/resources/locales/sl"
	"github.com/go-playground/universal-translator/resources/locales/smn"
	"github.com/go-playground/universal-translator/resources/locales/sn"
	"github.com/go-playground/universal-translator/resources/locales/so"
	"github.com/go-playground/universal-translator/resources/locales/so_DJ"
	"github.com/go-playground/universal-translator/resources/locales/so_ET"
	"github.com/go-playground/universal-translator/resources/locales/so_KE"
	"github.com/go-playground/universal-translator/resources/locales/sq"
	"github.com/go-playground/universal-translator/resources/locales/sq_MK"
	"github.com/go-playground/universal-translator/resources/locales/sr"
	"github.com/go-playground/universal-translator/resources/locales/sr_Latn"
	"github.com/go-playground/universal-translator/resources/locales/sv"
	"github.com/go-playground/universal-translator/resources/locales/sw"
	"github.com/go-playground/universal-translator/resources/locales/sw_CD"
	"github.com/go-playground/universal-translator/resources/locales/sw_UG"
	"github.com/go-playground/universal-translator/resources/locales/ta"
	"github.com/go-playground/universal-translator/resources/locales/ta_LK"
	"github.com/go-playground/universal-translator/resources/locales/ta_MY"
	"github.com/go-playground/universal-translator/resources/locales/ta_SG"
	"github.com/go-playground/universal-translator/resources/locales/te"
	"github.com/go-playground/universal-translator/resources/locales/teo"
	"github.com/go-playground/universal-translator/resources/locales/teo_KE"
	"github.com/go-playground/universal-translator/resources/locales/th"
	"github.com/go-playground/universal-translator/resources/locales/ti"
	"github.com/go-playground/universal-translator/resources/locales/ti_ER"
	"github.com/go-playground/universal-translator/resources/locales/tk"
	"github.com/go-playground/universal-translator/resources/locales/to"
	"github.com/go-playground/universal-translator/resources/locales/tr"
	"github.com/go-playground/universal-translator/resources/locales/twq"
	"github.com/go-playground/universal-translator/resources/locales/tzm"
	"github.com/go-playground/universal-translator/resources/locales/ug"
	"github.com/go-playground/universal-translator/resources/locales/uk"
	"github.com/go-playground/universal-translator/resources/locales/ur"
	"github.com/go-playground/universal-translator/resources/locales/ur_IN"
	"github.com/go-playground/universal-translator/resources/locales/uz"
	"github.com/go-playground/universal-translator/resources/locales/uz_Arab"
	"github.com/go-playground/universal-translator/resources/locales/uz_Cyrl"
	"github.com/go-playground/universal-translator/resources/locales/vai"
	"github.com/go-playground/universal-translator/resources/locales/vai_Latn"
	"github.com/go-playground/universal-translator/resources/locales/vi"
	"github.com/go-playground/universal-translator/resources/locales/vun"
	"github.com/go-playground/universal-translator/resources/locales/wae"
	"github.com/go-playground/universal-translator/resources/locales/xog"
	"github.com/go-playground/universal-translator/resources/locales/yav"
	"github.com/go-playground/universal-translator/resources/locales/yi"
	"github.com/go-playground/universal-translator/resources/locales/yo"
	"github.com/go-playground/universal-translator/resources/locales/yo_BJ"
	"github.com/go-playground/universal-translator/resources/locales/zgh"
	"github.com/go-playground/universal-translator/resources/locales/zh"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hans_HK"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hans_MO"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hans_SG"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hant"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hant_HK"
	"github.com/go-playground/universal-translator/resources/locales/zh_Hant_MO"
	"github.com/go-playground/universal-translator/resources/locales/zu"
)

func GetLocale(locale string) (*ut.Locale, error) {
	switch locale {
	case "af":
		return af.New(), nil
	case "af_NA":
		return af_NA.New(), nil
	case "agq":
		return agq.New(), nil
	case "ak":
		return ak.New(), nil
	case "am":
		return am.New(), nil
	case "ar":
		return ar.New(), nil
	case "ar_AE":
		return ar_AE.New(), nil
	case "ar_DJ":
		return ar_DJ.New(), nil
	case "ar_DZ":
		return ar_DZ.New(), nil
	case "ar_EH":
		return ar_EH.New(), nil
	case "ar_ER":
		return ar_ER.New(), nil
	case "ar_LB":
		return ar_LB.New(), nil
	case "ar_LY":
		return ar_LY.New(), nil
	case "ar_MA":
		return ar_MA.New(), nil
	case "ar_MR":
		return ar_MR.New(), nil
	case "ar_SO":
		return ar_SO.New(), nil
	case "ar_SS":
		return ar_SS.New(), nil
	case "ar_TN":
		return ar_TN.New(), nil
	case "as":
		return as.New(), nil
	case "asa":
		return asa.New(), nil
	case "ast":
		return ast.New(), nil
	case "az":
		return az.New(), nil
	case "az_Cyrl":
		return az_Cyrl.New(), nil
	case "bas":
		return bas.New(), nil
	case "be":
		return be.New(), nil
	case "bem":
		return bem.New(), nil
	case "bez":
		return bez.New(), nil
	case "bg":
		return bg.New(), nil
	case "bm":
		return bm.New(), nil
	case "bn":
		return bn.New(), nil
	case "bo":
		return bo.New(), nil
	case "bo_IN":
		return bo_IN.New(), nil
	case "br":
		return br.New(), nil
	case "brx":
		return brx.New(), nil
	case "bs":
		return bs.New(), nil
	case "bs_Cyrl":
		return bs_Cyrl.New(), nil
	case "ca":
		return ca.New(), nil
	case "ca_ES_VALENCIA":
		return ca_ES_VALENCIA.New(), nil
	case "ca_FR":
		return ca_FR.New(), nil
	case "ce":
		return ce.New(), nil
	case "cgg":
		return cgg.New(), nil
	case "chr":
		return chr.New(), nil
	case "ckb":
		return ckb.New(), nil
	case "cs":
		return cs.New(), nil
	case "cu":
		return cu.New(), nil
	case "cy":
		return cy.New(), nil
	case "da":
		return da.New(), nil
	case "dav":
		return dav.New(), nil
	case "de":
		return de.New(), nil
	case "de_AT":
		return de_AT.New(), nil
	case "de_CH":
		return de_CH.New(), nil
	case "de_LI":
		return de_LI.New(), nil
	case "de_LU":
		return de_LU.New(), nil
	case "dje":
		return dje.New(), nil
	case "dsb":
		return dsb.New(), nil
	case "dua":
		return dua.New(), nil
	case "dyo":
		return dyo.New(), nil
	case "dz":
		return dz.New(), nil
	case "ebu":
		return ebu.New(), nil
	case "ee":
		return ee.New(), nil
	case "el":
		return el.New(), nil
	case "en":
		return en.New(), nil
	case "en_001":
		return en_001.New(), nil
	case "en_150":
		return en_150.New(), nil
	case "en_AG":
		return en_AG.New(), nil
	case "en_AI":
		return en_AI.New(), nil
	case "en_AT":
		return en_AT.New(), nil
	case "en_AU":
		return en_AU.New(), nil
	case "en_BB":
		return en_BB.New(), nil
	case "en_BE":
		return en_BE.New(), nil
	case "en_BI":
		return en_BI.New(), nil
	case "en_BM":
		return en_BM.New(), nil
	case "en_BS":
		return en_BS.New(), nil
	case "en_BW":
		return en_BW.New(), nil
	case "en_BZ":
		return en_BZ.New(), nil
	case "en_CA":
		return en_CA.New(), nil
	case "en_CC":
		return en_CC.New(), nil
	case "en_CH":
		return en_CH.New(), nil
	case "en_CK":
		return en_CK.New(), nil
	case "en_CX":
		return en_CX.New(), nil
	case "en_DE":
		return en_DE.New(), nil
	case "en_DK":
		return en_DK.New(), nil
	case "en_DM":
		return en_DM.New(), nil
	case "en_ER":
		return en_ER.New(), nil
	case "en_FI":
		return en_FI.New(), nil
	case "en_FJ":
		return en_FJ.New(), nil
	case "en_FK":
		return en_FK.New(), nil
	case "en_GB":
		return en_GB.New(), nil
	case "en_GD":
		return en_GD.New(), nil
	case "en_GG":
		return en_GG.New(), nil
	case "en_GH":
		return en_GH.New(), nil
	case "en_GI":
		return en_GI.New(), nil
	case "en_GM":
		return en_GM.New(), nil
	case "en_GY":
		return en_GY.New(), nil
	case "en_IM":
		return en_IM.New(), nil
	case "en_IN":
		return en_IN.New(), nil
	case "en_JE":
		return en_JE.New(), nil
	case "en_JM":
		return en_JM.New(), nil
	case "en_KE":
		return en_KE.New(), nil
	case "en_KI":
		return en_KI.New(), nil
	case "en_KN":
		return en_KN.New(), nil
	case "en_KY":
		return en_KY.New(), nil
	case "en_LC":
		return en_LC.New(), nil
	case "en_LR":
		return en_LR.New(), nil
	case "en_LS":
		return en_LS.New(), nil
	case "en_MG":
		return en_MG.New(), nil
	case "en_MO":
		return en_MO.New(), nil
	case "en_MS":
		return en_MS.New(), nil
	case "en_MT":
		return en_MT.New(), nil
	case "en_MU":
		return en_MU.New(), nil
	case "en_MW":
		return en_MW.New(), nil
	case "en_MY":
		return en_MY.New(), nil
	case "en_NA":
		return en_NA.New(), nil
	case "en_NF":
		return en_NF.New(), nil
	case "en_NG":
		return en_NG.New(), nil
	case "en_NL":
		return en_NL.New(), nil
	case "en_NR":
		return en_NR.New(), nil
	case "en_NU":
		return en_NU.New(), nil
	case "en_NZ":
		return en_NZ.New(), nil
	case "en_PG":
		return en_PG.New(), nil
	case "en_PH":
		return en_PH.New(), nil
	case "en_PK":
		return en_PK.New(), nil
	case "en_PN":
		return en_PN.New(), nil
	case "en_RW":
		return en_RW.New(), nil
	case "en_SB":
		return en_SB.New(), nil
	case "en_SC":
		return en_SC.New(), nil
	case "en_SE":
		return en_SE.New(), nil
	case "en_SG":
		return en_SG.New(), nil
	case "en_SH":
		return en_SH.New(), nil
	case "en_SI":
		return en_SI.New(), nil
	case "en_SL":
		return en_SL.New(), nil
	case "en_SS":
		return en_SS.New(), nil
	case "en_SX":
		return en_SX.New(), nil
	case "en_SZ":
		return en_SZ.New(), nil
	case "en_TK":
		return en_TK.New(), nil
	case "en_TO":
		return en_TO.New(), nil
	case "en_TT":
		return en_TT.New(), nil
	case "en_TV":
		return en_TV.New(), nil
	case "en_TZ":
		return en_TZ.New(), nil
	case "en_UG":
		return en_UG.New(), nil
	case "en_US_POSIX":
		return en_US_POSIX.New(), nil
	case "en_VC":
		return en_VC.New(), nil
	case "en_VU":
		return en_VU.New(), nil
	case "en_WS":
		return en_WS.New(), nil
	case "en_ZA":
		return en_ZA.New(), nil
	case "en_ZM":
		return en_ZM.New(), nil
	case "eo":
		return eo.New(), nil
	case "es":
		return es.New(), nil
	case "es_419":
		return es_419.New(), nil
	case "es_AR":
		return es_AR.New(), nil
	case "es_BO":
		return es_BO.New(), nil
	case "es_CL":
		return es_CL.New(), nil
	case "es_CO":
		return es_CO.New(), nil
	case "es_CR":
		return es_CR.New(), nil
	case "es_CU":
		return es_CU.New(), nil
	case "es_DO":
		return es_DO.New(), nil
	case "es_EC":
		return es_EC.New(), nil
	case "es_GQ":
		return es_GQ.New(), nil
	case "es_GT":
		return es_GT.New(), nil
	case "es_HN":
		return es_HN.New(), nil
	case "es_MX":
		return es_MX.New(), nil
	case "es_NI":
		return es_NI.New(), nil
	case "es_PA":
		return es_PA.New(), nil
	case "es_PE":
		return es_PE.New(), nil
	case "es_PH":
		return es_PH.New(), nil
	case "es_PR":
		return es_PR.New(), nil
	case "es_PY":
		return es_PY.New(), nil
	case "es_SV":
		return es_SV.New(), nil
	case "es_US":
		return es_US.New(), nil
	case "es_UY":
		return es_UY.New(), nil
	case "es_VE":
		return es_VE.New(), nil
	case "et":
		return et.New(), nil
	case "eu":
		return eu.New(), nil
	case "ewo":
		return ewo.New(), nil
	case "fa":
		return fa.New(), nil
	case "fa_AF":
		return fa_AF.New(), nil
	case "ff":
		return ff.New(), nil
	case "ff_GN":
		return ff_GN.New(), nil
	case "ff_MR":
		return ff_MR.New(), nil
	case "fi":
		return fi.New(), nil
	case "fil":
		return fil.New(), nil
	case "fo":
		return fo.New(), nil
	case "fo_DK":
		return fo_DK.New(), nil
	case "fr":
		return fr.New(), nil
	case "fr_BE":
		return fr_BE.New(), nil
	case "fr_BI":
		return fr_BI.New(), nil
	case "fr_CA":
		return fr_CA.New(), nil
	case "fr_CD":
		return fr_CD.New(), nil
	case "fr_CH":
		return fr_CH.New(), nil
	case "fr_DJ":
		return fr_DJ.New(), nil
	case "fr_DZ":
		return fr_DZ.New(), nil
	case "fr_GN":
		return fr_GN.New(), nil
	case "fr_HT":
		return fr_HT.New(), nil
	case "fr_KM":
		return fr_KM.New(), nil
	case "fr_LU":
		return fr_LU.New(), nil
	case "fr_MA":
		return fr_MA.New(), nil
	case "fr_MG":
		return fr_MG.New(), nil
	case "fr_MR":
		return fr_MR.New(), nil
	case "fr_MU":
		return fr_MU.New(), nil
	case "fr_RW":
		return fr_RW.New(), nil
	case "fr_SC":
		return fr_SC.New(), nil
	case "fr_SY":
		return fr_SY.New(), nil
	case "fr_TN":
		return fr_TN.New(), nil
	case "fr_VU":
		return fr_VU.New(), nil
	case "fur":
		return fur.New(), nil
	case "fy":
		return fy.New(), nil
	case "ga":
		return ga.New(), nil
	case "gd":
		return gd.New(), nil
	case "gl":
		return gl.New(), nil
	case "gsw":
		return gsw.New(), nil
	case "gu":
		return gu.New(), nil
	case "guz":
		return guz.New(), nil
	case "gv":
		return gv.New(), nil
	case "ha":
		return ha.New(), nil
	case "ha_GH":
		return ha_GH.New(), nil
	case "haw":
		return haw.New(), nil
	case "he":
		return he.New(), nil
	case "hi":
		return hi.New(), nil
	case "hr":
		return hr.New(), nil
	case "hr_BA":
		return hr_BA.New(), nil
	case "hsb":
		return hsb.New(), nil
	case "hu":
		return hu.New(), nil
	case "hy":
		return hy.New(), nil
	case "id":
		return id.New(), nil
	case "ig":
		return ig.New(), nil
	case "ii":
		return ii.New(), nil
	case "is":
		return is.New(), nil
	case "it":
		return it.New(), nil
	case "it_CH":
		return it_CH.New(), nil
	case "ja":
		return ja.New(), nil
	case "jgo":
		return jgo.New(), nil
	case "jmc":
		return jmc.New(), nil
	case "ka":
		return ka.New(), nil
	case "kab":
		return kab.New(), nil
	case "kam":
		return kam.New(), nil
	case "kde":
		return kde.New(), nil
	case "kea":
		return kea.New(), nil
	case "khq":
		return khq.New(), nil
	case "ki":
		return ki.New(), nil
	case "kk":
		return kk.New(), nil
	case "kkj":
		return kkj.New(), nil
	case "kl":
		return kl.New(), nil
	case "kln":
		return kln.New(), nil
	case "km":
		return km.New(), nil
	case "kn":
		return kn.New(), nil
	case "ko":
		return ko.New(), nil
	case "kok":
		return kok.New(), nil
	case "ks":
		return ks.New(), nil
	case "ksb":
		return ksb.New(), nil
	case "ksf":
		return ksf.New(), nil
	case "ksh":
		return ksh.New(), nil
	case "kw":
		return kw.New(), nil
	case "ky":
		return ky.New(), nil
	case "lag":
		return lag.New(), nil
	case "lb":
		return lb.New(), nil
	case "lg":
		return lg.New(), nil
	case "lkt":
		return lkt.New(), nil
	case "ln":
		return ln.New(), nil
	case "ln_AO":
		return ln_AO.New(), nil
	case "lo":
		return lo.New(), nil
	case "lrc":
		return lrc.New(), nil
	case "lt":
		return lt.New(), nil
	case "lu":
		return lu.New(), nil
	case "luo":
		return luo.New(), nil
	case "luy":
		return luy.New(), nil
	case "lv":
		return lv.New(), nil
	case "mas":
		return mas.New(), nil
	case "mas_TZ":
		return mas_TZ.New(), nil
	case "mer":
		return mer.New(), nil
	case "mfe":
		return mfe.New(), nil
	case "mg":
		return mg.New(), nil
	case "mgh":
		return mgh.New(), nil
	case "mgo":
		return mgo.New(), nil
	case "mk":
		return mk.New(), nil
	case "ml":
		return ml.New(), nil
	case "mn":
		return mn.New(), nil
	case "mr":
		return mr.New(), nil
	case "ms":
		return ms.New(), nil
	case "ms_BN":
		return ms_BN.New(), nil
	case "ms_SG":
		return ms_SG.New(), nil
	case "mt":
		return mt.New(), nil
	case "mua":
		return mua.New(), nil
	case "my":
		return my.New(), nil
	case "mzn":
		return mzn.New(), nil
	case "naq":
		return naq.New(), nil
	case "nb":
		return nb.New(), nil
	case "nd":
		return nd.New(), nil
	case "ne":
		return ne.New(), nil
	case "nl":
		return nl.New(), nil
	case "nl_AW":
		return nl_AW.New(), nil
	case "nl_BE":
		return nl_BE.New(), nil
	case "nl_BQ":
		return nl_BQ.New(), nil
	case "nl_CW":
		return nl_CW.New(), nil
	case "nl_SR":
		return nl_SR.New(), nil
	case "nl_SX":
		return nl_SX.New(), nil
	case "nmg":
		return nmg.New(), nil
	case "nn":
		return nn.New(), nil
	case "nnh":
		return nnh.New(), nil
	case "nus":
		return nus.New(), nil
	case "nyn":
		return nyn.New(), nil
	case "om":
		return om.New(), nil
	case "om_KE":
		return om_KE.New(), nil
	case "or":
		return or.New(), nil
	case "os":
		return os.New(), nil
	case "os_RU":
		return os_RU.New(), nil
	case "pa":
		return pa.New(), nil
	case "pa_Arab":
		return pa_Arab.New(), nil
	case "pl":
		return pl.New(), nil
	case "prg":
		return prg.New(), nil
	case "ps":
		return ps.New(), nil
	case "pt":
		return pt.New(), nil
	case "pt_AO":
		return pt_AO.New(), nil
	case "pt_CV":
		return pt_CV.New(), nil
	case "pt_MO":
		return pt_MO.New(), nil
	case "pt_MZ":
		return pt_MZ.New(), nil
	case "pt_PT":
		return pt_PT.New(), nil
	case "pt_ST":
		return pt_ST.New(), nil
	case "qu":
		return qu.New(), nil
	case "qu_BO":
		return qu_BO.New(), nil
	case "qu_EC":
		return qu_EC.New(), nil
	case "rm":
		return rm.New(), nil
	case "rn":
		return rn.New(), nil
	case "ro":
		return ro.New(), nil
	case "ro_MD":
		return ro_MD.New(), nil
	case "rof":
		return rof.New(), nil
	case "root":
		return root.New(), nil
	case "ru":
		return ru.New(), nil
	case "ru_BY":
		return ru_BY.New(), nil
	case "ru_KG":
		return ru_KG.New(), nil
	case "ru_KZ":
		return ru_KZ.New(), nil
	case "ru_MD":
		return ru_MD.New(), nil
	case "rw":
		return rw.New(), nil
	case "rwk":
		return rwk.New(), nil
	case "sah":
		return sah.New(), nil
	case "saq":
		return saq.New(), nil
	case "sbp":
		return sbp.New(), nil
	case "se":
		return se.New(), nil
	case "se_SE":
		return se_SE.New(), nil
	case "seh":
		return seh.New(), nil
	case "ses":
		return ses.New(), nil
	case "sg":
		return sg.New(), nil
	case "shi":
		return shi.New(), nil
	case "shi_Latn":
		return shi_Latn.New(), nil
	case "si":
		return si.New(), nil
	case "sk":
		return sk.New(), nil
	case "sl":
		return sl.New(), nil
	case "smn":
		return smn.New(), nil
	case "sn":
		return sn.New(), nil
	case "so":
		return so.New(), nil
	case "so_DJ":
		return so_DJ.New(), nil
	case "so_ET":
		return so_ET.New(), nil
	case "so_KE":
		return so_KE.New(), nil
	case "sq":
		return sq.New(), nil
	case "sq_MK":
		return sq_MK.New(), nil
	case "sr":
		return sr.New(), nil
	case "sr_Latn":
		return sr_Latn.New(), nil
	case "sv":
		return sv.New(), nil
	case "sw":
		return sw.New(), nil
	case "sw_CD":
		return sw_CD.New(), nil
	case "sw_UG":
		return sw_UG.New(), nil
	case "ta":
		return ta.New(), nil
	case "ta_LK":
		return ta_LK.New(), nil
	case "ta_MY":
		return ta_MY.New(), nil
	case "ta_SG":
		return ta_SG.New(), nil
	case "te":
		return te.New(), nil
	case "teo":
		return teo.New(), nil
	case "teo_KE":
		return teo_KE.New(), nil
	case "th":
		return th.New(), nil
	case "ti":
		return ti.New(), nil
	case "ti_ER":
		return ti_ER.New(), nil
	case "tk":
		return tk.New(), nil
	case "to":
		return to.New(), nil
	case "tr":
		return tr.New(), nil
	case "twq":
		return twq.New(), nil
	case "tzm":
		return tzm.New(), nil
	case "ug":
		return ug.New(), nil
	case "uk":
		return uk.New(), nil
	case "ur":
		return ur.New(), nil
	case "ur_IN":
		return ur_IN.New(), nil
	case "uz":
		return uz.New(), nil
	case "uz_Arab":
		return uz_Arab.New(), nil
	case "uz_Cyrl":
		return uz_Cyrl.New(), nil
	case "vai":
		return vai.New(), nil
	case "vai_Latn":
		return vai_Latn.New(), nil
	case "vi":
		return vi.New(), nil
	case "vun":
		return vun.New(), nil
	case "wae":
		return wae.New(), nil
	case "xog":
		return xog.New(), nil
	case "yav":
		return yav.New(), nil
	case "yi":
		return yi.New(), nil
	case "yo":
		return yo.New(), nil
	case "yo_BJ":
		return yo_BJ.New(), nil
	case "zgh":
		return zgh.New(), nil
	case "zh":
		return zh.New(), nil
	case "zh_Hans_HK":
		return zh_Hans_HK.New(), nil
	case "zh_Hans_MO":
		return zh_Hans_MO.New(), nil
	case "zh_Hans_SG":
		return zh_Hans_SG.New(), nil
	case "zh_Hant":
		return zh_Hant.New(), nil
	case "zh_Hant_HK":
		return zh_Hant_HK.New(), nil
	case "zh_Hant_MO":
		return zh_Hant_MO.New(), nil
	case "zu":
		return zu.New(), nil

	default:
		return nil, errors.New("Unknown locale '" + locale + "'")
	}
}
