package config

var QueueName = "queueIndex"
var VisitedIndexName = "visitedIndex"

// AcceptableMimeTypes Mimes aceitos, checagem quando visitado
var AcceptableMimeTypes = []string{
	"text/html",
	"text/plain",
	"text/xml",
	"application/xml",
	"application/xhtml+xml",
	"application/rss+xml",
	"application/atom+xml",
	"application/rdf+xml",
	"application/json",
	"application/ld+json",
	"application/vnd.geo+json",
	"application/xml-dtd",
	"application/rss+xml",
	"application/atom+xml",
	"application/rdf+xml",
	"application/json",
	"application/ld+json",
	"application/vnd.geo+json",
}

// AcceptableSchema Schemas permitidos
var AcceptableSchema = []string{
	"http",
	"https",
	"",
}

// DenySuffixes Impede urls com estes sufixos de serem visitadas.
var DenySuffixes = []string{
	".css",
	".js",
	".png",
	".jpg",
	".jpeg",
	".gif",
	".svg",
	".ico",
	".mp4",
	".mp3",
	".avi",
	".flv",
	".mpeg",
	".webp",
	".webm",
	".woff",
	".woff2",
	".ttf",
	".eot",
	".otf",
	".pdf",
	".zip",
	".tar",
	".gz",
	".bz2",
	".xz",
	".7z",
	".rar",
	".apk",
	".exe",
	".dmg",
	".img",
	".pdf",
}

// CommonStopWords Palavras de parada comuns (personalize conforme necessário, palavras geradas por GPT)
var CommonStopWords = map[string][]string{
	"en": {"is", "or", "a", "and", "the", "are", "of", "to"},
	"pt": {
		"a", "à", "ao", "aos", "aquela", "aquelas", "aquele", "aqueles", "aquilo", "as", "até",
		"com", "como", "da", "das", "de", "dela", "delas", "dele", "deles", "desta", "destas",
		"deste", "destes", "do", "dos", "e", "é", "ela", "elas", "ele", "eles", "em", "entre",
		"era", "essa", "essas", "esse", "esses", "esta", "estamos", "estas", "estava", "estavam",
		"estávamos", "este", "estes", "estou", "eu", "foi", "fomos", "for", "foram", "havia",
		"isso", "isto", "já", "lhe", "lhes", "mais", "mas", "me", "mesmo", "meu", "meus", "minha",
		"minhas", "muito", "na", "nas", "nem", "no", "nos", "nossa", "nossas", "nosso", "nossos",
		"num", "numa", "o", "os", "ou", "para", "pela", "pelas", "pelo", "pelos", "por", "qual",
		"quando", "que", "quem", "se", "sem", "seu", "seus", "sua", "suas", "só", "também", "te",
		"tem", "temos", "tenha", "tenham", "teu", "teus", "teve", "tinha", "tinham", "tua", "tuas",
		"um", "uma", "você", "vocês", "vos",
	},
	"ru": {
		"а", "без", "более", "бы", "был", "была", "были", "было", "быть", "в", "вам", "вас", "всё",
		"все", "всего", "всех", "вы", "где", "да", "даже", "для", "до", "его", "ее", "ей", "ему",
		"если", "ест", "есть", "ещё", "ж", "же", "за", "здесь", "и", "из", "или", "им", "их", "к",
		"как", "какая", "какой", "когда", "кое", "кто", "куда", "ли", "либо", "мне", "много", "может",
		"можно", "мой", "моя", "мы", "на", "над", "надо", "наш", "не", "него", "нее", "нет", "ни",
		"них", "но", "ну", "о", "об", "однако", "он", "она", "они", "оно", "от", "очень", "по",
		"под", "после", "при", "про", "с", "сам", "сама", "сами", "само", "свое", "своего", "своей",
		"свои", "себе", "себя", "сейчас", "со", "совсем", "так", "такой", "там", "тебя", "тем",
		"теперь", "то", "тогда", "того", "тоже", "только", "том", "ты", "у", "уже", "хотя", "чего",
		"чей", "чем", "что", "чтобы", "чуть", "эта", "эти", "это", "этого", "этой", "этом", "эту",
		"я",
	},
	"es": {
		"el", "la", "los", "las", "de", "del", "y", "a", "en", "un", "una", "unos", "unas", "con",
		"para", "por", "su", "se", "que", "es", "soy", "eres", "somos", "son", "me", "te", "nos", "le",
		"les", "lo", "mi", "tu", "si", "no", "pero", "porque", "como", "esta", "estoy", "estas", "estamos",
		"estais", "estan", "muy", "poco", "mucho", "todo", "todos", "al", "algo", "alguien", "donde", "cuando",
		"como", "aqui", "ahi", "alli", "ahora", "antes", "despues", "hoy", "ayer", "mañana", "siempre", "nunca",
	},
	"hindi": {
		"का", "के", "की", "में", "है", "और", "यह", "वह", "से", "को", "पर", "इस", "होता", "ही", "हैं", "ये", "वो", "कर", "गया", "लिए",
		"अपना", "अपनी", "अपने", "कुछ", "थी", "थे", "थीं", "हुआ", "जा", "रहा", "रहे", "जाता", "जाती", "जाते", "एक", "दो", "तीन", "चार",
		"पांच", "छह", "सात", "आठ", "नौ", "दस",
	},
	"ch": {
		"的", "了", "在", "是", "我", "有", "和", "就", "不", "人", "这", "那", "中", "来", "上", "大", "为", "个", "国",
		"以", "说", "到", "要", "子", "你", "会", "着", "能", "里", "去", "年", "得", "他", "她", "它", "们", "地", "也",
		"自", "这", "时", "那", "儿", "可", "就", "给", "下", "都", "向", "看", "起", "还", "过", "只", "把", "对", "做",
		"当", "想", "成", "事", "被", "用", "多", "从", "面", "等", "前", "些", "于", "后", "所", "又", "经", "方", "现",
		"没", "吧", "定", "得", "该", "好好", "家", "种", "那", "里", "然", "其", "间", "什", "么", "很", "得", "哪", "些",
		"向", "生", "里", "果", "再", "两", "并", "而", "些", "定",
	},
}
