package stopwords

var BrazilianPortuguese = []string{
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
	"um", "uma", "você", "vocês", "vos", "agora", "ainda", "ano", "anos", "antes", "aonde",
	"apenas", "apoio", "apontar", "aqui", "assim", "atrás", "baixo", "bem", "boa", "boas",
	"bom", "bons", "breve", "cada", "caminho", "cedo", "cima", "coisa", "comprido", "conhecido",
	"conselho", "contra", "corrente", "custa", "cá", "dar", "debaixo", "demais", "dentro",
	"depois", "desde", "desligado", "deve", "devem", "dia", "diante", "direita", "diversa",
	"diversas", "diversos", "diz", "dizem", "dizer", "dois", "duas", "durante", "dá", "dão",
	"embora", "enquanto", "então", "eram", "estado", "estar", "estará", "esteja", "estejam",
	"estejamos", "esteve", "estive", "estivemos", "estiver", "estivera", "estiveram", "estiverem",
	"estivermos", "estivesse", "estivessem", "estiveste", "estivestes", "estivéramos", "estivéssemos",
	"exemplo", "falta", "fará", "favor", "faz", "fazeis", "fazem", "fazemos", "fazer", "fazes",
	"fazia", "faço", "fez", "fim", "final", "fora", "forem", "forma", "formos", "fosse",
	"fossem", "foste", "fostes", "fui", "fôramos", "fôssemos", "geral", "grande", "grandes",
	"grupo", "ha", "haja", "hajam", "hajamos", "havemos", "hei", "hoje", "hora", "horas",
	"houve", "houvemos", "houver", "houvera", "houveram", "houverei", "houverem", "houveremos",
	"houveria", "houveriam", "houvermos", "houverá", "houverão", "houveríamos", "houvesse",
	"houvessem", "houvéramos", "houvéssemos", "há", "hão", "iniciar", "inicio", "ir", "irá",
	"ista", "iste", "lado", "ligado", "local", "logo", "longe", "lugar", "lá", "maior",
	"maioria", "maiorias", "mal", "mediante", "meio", "menor", "menos", "meses", "mesma",
	"mesmas", "mesmos", "mil", "momento", "muitos", "máximo", "mês", "nada", "nao", "naquela",
	"naquelas", "naquele", "naqueles", "nenhuma", "nessa", "nessas", "nesse", "nesses", "nesta",
	"nestas", "neste", "nestes", "noite", "nome", "nova", "novas", "nove", "novo", "novos",
	"numas", "nunca", "nuns", "não", "nível", "nós", "número", "obra", "obrigada", "obrigado",
	"oitava", "oitavo", "oito", "onde", "ontem", "onze", "outra", "outras", "outro", "outros",
	"parece", "parte", "partir", "pegar", "perante", "perto", "pessoas", "pode", "podem", "poder",
	"poderá", "podia", "pois", "ponto", "pontos", "porque", "porquê", "portanto", "posição",
	"possivelmente", "posso", "possível", "pouca", "pouco", "poucos", "povo", "primeira",
	"primeiras", "primeiro", "primeiros", "propios", "proprio", "própria", "próprias",
	"próprio", "próprios", "próxima", "próximas", "próximo", "próximos", "puderam", "pôde",
	"põe", "põem", "quais", "qualquer", "quanto", "quarta", "quarto", "quatro", "quer",
	"quereis", "querem", "queremas", "queres", "quero", "questão", "quieto", "quinta",
	"quinto", "quinze", "quáis", "quê", "relação", "sabe", "sabem", "saber", "segunda",
	"segundo", "sei", "seis", "seja", "sejam", "sejamos", "sempre", "sendo", "ser", "serei",
	"seremos", "seria", "seriam", "será", "serão", "seríamos", "sete", "sexta", "sexto", "sim",
	"sistema", "sob", "sobre", "sois", "somente", "somos", "sou", "são", "sétima", "sétimo",
	"tal", "talvez", "tambem", "tanta", "tantas", "tanto", "tarde", "tempo", "tendes", "tenhamos",
	"tenho", "tens", "tentar", "tentaram", "tente", "tentei", "ter", "terceira", "terceiro",
	"terei", "teremos", "teria", "teriam", "terá", "terão", "teríamos", "tipo", "tive", "tivemos",
	"tiverem", "tivermos", "tivesse", "tivessem", "tiveste", "tivestes", "tivéramos", "tivéssemos",
	"toda", "todas", "todo", "todos", "trabalhar", "trabalho", "treze", "três", "tu", "tudo",
	"tão", "tém", "têm", "tínhamos", "umas", "uns", "usa", "usar", "vai", "vais", "valor",
	"veja", "vem", "vens", "ver", "verdade", "verdadeiro", "vez", "vezes", "viagem", "vindo",
	"vinte", "vossa", "vossas", "vossos", "vários", "vão", "vêm", "vós", "zero", "às", "área",
	"éramos", "és", "último",
}

var EuropeanPortuguese = []string{
	"a", "à", "acerca", "adeus", "ainda", "alem", "algmas", "algo", "algumas", "alguns", "ali",
	"além", "ambas", "ambos", "ao", "aonde", "aos", "apenas", "apos", "aquela", "aquelas",
	"aquele", "aqueles", "aqui", "aquilo", "as", "assim", "através", "até", "aí", "baixo",
	"bastante", "bem", "catorze", "cento", "certamente", "certeza", "cinco", "com", "como",
	"contra", "contudo", "cuja", "cujas", "cujo", "cujos", "da", "daquela", "daquelas",
	"daquele", "daqueles", "dar", "das", "de", "dela", "delas", "dele", "deles", "demais",
	"dentro", "depois", "desde", "dessa", "dessas", "desse", "desses", "desta", "destas",
	"deste", "destes", "deve", "devem", "deverá", "dez", "dezanove", "dezasseis", "dezassete",
	"dezoito", "dia", "diante", "dispoe", "dispoem", "diversa", "diversas", "diversos", "diz",
	"dizem", "dizer", "do", "dois", "dos", "doze", "duas", "durante", "dá", "dão", "e", "ela",
	"elas", "ele", "eles", "em", "enquanto", "entao", "entre", "então", "era", "eram", "essa",
	"essas", "esse", "esses", "esta", "estado", "estamos", "estar", "estará", "estas", "estava",
	"estavam", "este", "esteja", "estejam", "estejamos", "estes", "esteve", "estive",
	"estivemos", "estiver", "estivera", "estiveram", "estiverem", "estivermos", "estivesse",
	"estivessem", "estiveste", "estivestes", "estivéramos", "estivéssemos", "estou", "está",
	"estás", "estávamos", "estão", "eu", "exemplo", "falta", "fará", "favor", "faz", "fazeis",
	"fazem", "fazemos", "fazer", "fazes", "fazia", "faço", "fez", "fim", "final", "foi", "fomos",
	"for", "fora", "foram", "forem", "forma", "formos", "fosse", "fossem", "foste", "fostes",
	"fui", "fôramos", "fôssemos", "geral", "grande", "grandes", "grupo", "haja", "hajam",
	"hajamos", "havemos", "havia", "hei", "hoje", "hora", "horas", "houve", "houvemos",
	"houver", "houvera", "houveram", "houverei", "houverem", "houveremos", "houveria",
	"houveriam", "houvermos", "houverá", "houverão", "houveríamos", "houvesse", "houvessem",
	"houvéramos", "houvéssemos", "há", "hão", "iniciar", "inicio", "ir", "irá", "isso", "isto",
	"já", "lado", "lhe", "lhes", "ligado", "local", "logo", "longe", "lugar", "lá", "maior",
	"maioria", "maiorias", "mais", "mal", "mas", "me", "mediante", "meio", "menor", "menos",
	"meses", "mesma", "mesmas", "mesmo", "mesmos", "meu", "meus", "mil", "minha", "minhas",
	"momento", "muito", "muitos", "máximo", "mês", "na", "nada", "nao", "naquela", "naquelas",
	"naquele", "naqueles", "nas", "nem", "nenhuma", "nessa", "nessas", "nesse", "nesses",
	"nesta", "nestas", "neste", "nestes", "no", "noite", "nome", "nos", "nossa", "nossas",
	"nosso", "nossos", "nova", "novas", "nove", "novo", "novos", "num", "numa", "numas",
	"nunca", "nuns", "não", "nível", "nós", "número", "obra", "obrigada", "obrigado", "oitava",
	"oitavo", "oito", "onde", "ontem", "onze", "os", "ou", "outra", "outras", "outro", "outros",
	"para", "parece", "parte", "partir", "paucas", "pela", "pelas", "pelo", "pelos", "perante",
	"perto", "pessoas", "pode", "podem", "poder", "poderá", "podia", "pois", "ponto", "pontos",
	"por", "porque", "porquê", "portanto", "posição", "possivelmente", "posso", "possível",
	"pouca", "pouco", "poucos", "povo", "primeira", "primeiras", "primeiro", "primeiros",
	"promeiro", "própria", "próprias", "próprio", "próprios", "próxima", "próximas", "próximo",
	"próximos", "puderam", "pôde", "põe", "põem", "quais", "qual", "qualquer", "quando", "quanto",
	"quarta", "quarto", "quatro", "que", "quem", "quer", "quereis", "querem", "queremas",
	"queres", "quero", "questão", "quieto", "quinta", "quinto", "quinze", "quáis", "quê",
	"relação", "sabe", "sabem", "saber", "se", "segunda", "segundo", "sei", "seis", "seja",
	"sejam", "sejamos", "sem", "sempre", "sendo", "ser", "serei", "seremos", "seria", "seriam",
	"será", "serão", "seríamos", "sete", "seu", "seus", "sexta", "sexto", "sim", "sistema",
	"sob", "sobre", "sois", "somente", "somos", "sou", "sua", "suas", "são", "sétima", "sétimo",
	"só", "tal", "talvez", "também", "tanta", "tantas", "tanto", "tarde", "te", "tem", "temos",
	"tempo", "tendes", "tenha", "tenham", "tenhamos", "tenho", "tens", "tentar", "tentaram",
	"tente", "tentei", "ter", "terceira", "terceiro", "terei", "teremos", "teria", "teriam",
	"terá", "terão", "teríamos", "teu", "teus", "teve", "tinha", "tinham", "tive", "tivemos",
	"tiver", "tivera", "tiveram", "tiverem", "tivermos", "tivesse", "tivessem", "tiveste",
	"tivestes", "tivéramos", "tivéssemos", "toda", "todas", "todo", "todos", "trabalhar",
	"trabalho", "treze", "três", "tu", "tua", "tuas", "tudo", "tão", "tém", "têm", "tínhamos",
	"um", "uma", "umas", "uns", "usa", "usar", "vai", "vais", "valor", "veja", "vem", "vens",
	"ver", "verdade", "verdadeiro", "vez", "vezes", "viagem", "vindo", "vinte", "você", "vocês",
	"vos", "vossa", "vossas", "vosso", "vossos", "vários", "vão", "vêm", "vós", "zero",
}