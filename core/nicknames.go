package core

import (
	"encoding/json"
)

func FindNickNames(name string) (bool, []string) {
	result := []string{}
	names := map[string][]string{}

	json.Unmarshal([]byte(NAME_TO_NICKNAME), &names)
	value, exists := names[name]
	if exists {
		return exists, value
	}

	return false, result
}

func FindNameFromNickName(name string) (bool, string) {
	var result string
	names := map[string]string{}

	json.Unmarshal([]byte(Nickname_To_Name), &names)
	value, exists := names[name]
	if exists {
		return exists, value
	}

	return false, result
}

const NAME_TO_NICKNAME = `{
    "abigail": ["abbey", "abbi", "abbie", "abby", "abi", "gail", "gayle"],
    "adelaide": ["addie"],
    "adele": ["addie"],
    "adeline": ["addie"],
    "agatha": ["aggie"],
    "agnes": ["aggie", "nessie"],
    "alex": ["al", "lex"],
    "alexa": ["lexi"],
    "alexander": ["al", "alec", "alex", "aleck", "elsie", "lex", "sander", "sandy", "sawney", "xan", "zander"],
    "alexandra": ["allie", "ally", "aly", "lexi", "lexie", "sandy"],
    "alexia": ["lexi", "lexie"],
    "alexis": ["lexi", "lexie"],
    "alfonso": ["alf", "alfie", "alfy"],
    "alfonzo": ["alf", "alfie", "alfy", "fonz", "fonzie"],
    "alfred": ["al", "alf", "alfie", "alfy", "allie", "ally", "fred"],
    "alice": ["ally"],
    "alison": ["ali", "allie", "ally", "aly"],
    "alistair": ["allie"],
    "allie": ["ali"],
    "allison": ["ali"],
    "allyson": ["ali"],
    "alonzo": ["lonnie", "lonny"],
    "alvin": ["allie", "alvie", "alvy"],
    "alyson": ["ali"],
    "alyssa": ["aly"],
    "amanda": ["mandy"],
    "ambrose": ["ambers", "ambie", "amby", "brose"],
    "amelia": ["millie"],
    "andrea": ["andy"],
    "andrew": ["andy", "drew"],
    "angela": ["ange", "angie"],
    "angelika": ["ange", "angie"],
    "angelina": ["ange", "angie"],
    "anjana": ["anj"],
    "ann": ["annie", "jo", "nan", "nannie", "nanny"],
    "annabelle": ["belle"],
    "anne": ["nanny"],
    "annette": ["nettie"],
    "annie": ["nannie"],
    "anthony": ["ant", "tony"],
    "antonia": ["toni", "tonia"],
    "antoinette": ["toni"],
    "arabella": ["bella"],
    "archibald": ["archie"],
    "arnold": ["arnie"],
    "arthur": ["arfie", "art", "artie", "arty"],
    "ashlee": ["ash"],
    "ashley": ["ash"],
    "audrey": ["audie"],
    "augustus": ["augie", "gus"],
    "barbara": ["babs", "barb", "barbie"],
    "barnabas": ["barney"],
    "barnaby": ["barney"],
    "barnett": ["barney"],
    "barry": ["baz", "bazza"],
    "bartholomew": ["bart", "batt"],
    "basil": ["baz", "bazza"],
    "beatrice": ["bea", "trixie"],
    "beatrix": ["trixie"],
    "becky": ["becki", "beki"],
    "belle": ["bell"],
    "benedict": ["ben"],
    "benjamin": ["ben", "benji", "benjie", "benjy", "bennie", "benny"],
    "bernadette": ["bennie", "benny"],
    "bernard": ["barney", "bernie"],
    "bernice": ["bennie", "benny"],
    "bert": ["bertie"],
    "bertram": ["bertie"],
    "bessie": ["bessy"],
    "bethany": ["beth"],
    "bethwyn": ["beth"],
    "betty": ["bette"],
    "bradford": ["brad"],
    "bradley": ["brad"],
    "bradly": ["brad"],
    "brian": ["bri"],
    "bridget": ["biddie", "biddy"],
    "caden": ["cade"],
    "caleb": ["cal", "cale"],
    "calum": ["cal"],
    "calvin": ["cal"],
    "cameron": ["cam"],
    "camilla": ["millie", "milly"],
    "campbell": ["camp"],
    "candace": ["candi", "candy"],
    "candice": ["candy"],
    "caroline": ["cal", "callie", "caro", "carrie", "cary"],
    "carolyn": ["lyn"],
    "cassandra": ["cass", "cassie"],
    "catherine": ["cat", "cate", "cath", "cathi", "cathie", "cathy", "catie", "caty", "kate", "katie"],
    "cathy": ["cathi", "cathie"],
    "cecilia": ["ceecee", "ceecee", "sissy"],
    "cecily": ["cecie"],
    "chadwick": ["chad"],
    "charlene": ["charley", "charlie"],
    "charles": ["chad", "charley", "charlie", "chas", "chaz", "chip", "chuck"],
    "charlotte": ["charley", "charlie", "lottie"],
    "chauncey": ["chance"],
    "cherilyn": ["cher"],
    "chester": ["chet"],
    "christian": ["chris", "kris"],
    "christie": ["christy", "kristi", "kristie", "kristy"],
    "christina": ["chris", "chrissie", "chrissy", "christy", "kris", "krissy", "kristi", "kristie", "kristy"],
    "christine": ["chris", "chrissie", "chrissy", "christie", "krissy"],
    "christopher": ["chip", "chris", "christie", "kit", "kris"],
    "chuck": ["chuckie", "chucky"],
    "clayton": ["clay"],
    "clement": ["clem", "clemmie"],
    "clementine": ["clem"],
    "clifford": ["cliff"],
    "clinton": ["clint"],
    "conor": ["con"],
    "constance": ["connie"],
    "cora": ["corrie"],
    "corinne": ["corrie"],
    "cornelia": ["corrie"],
    "cornelius": ["con"],
    "curtis": ["curt"],
    "cynthia": ["cindy"],
    "cyril": ["cy"],
    "cyrus": ["cy"],
    "daniel": ["dan", "dani", "danny"],
    "daniela": ["danny"],
    "daniella": ["danny"],
    "danielle": ["dani", "danni", "danny"],
    "danny": ["dan", "dani", "dannie"],
    "david": ["dave", "davey", "davie", "davo", "davy"],
    "deborah": ["deb", "debbie", "debby"],
    "debra": ["deb", "debbie", "debby"],
    "denis": ["denny"],
    "dennis": ["den", "denny"],
    "derek": ["del"],
    "dermot": ["jeremy"],
    "desmond": ["des"],
    "diana": ["di"],
    "diane": ["di"],
    "diarmaid": ["dermot", "jeremy"],
    "diarmid": ["diarmi", "dermot", "jeremy"],
    "diarmuid": ["dermot", "jeremy"],
    "dick": ["dickie"],
    "dolores": ["dolly"],
    "dominic": ["dom"],
    "don": ["donnie", "donny"],
    "donald": ["don", "donnie", "donny"],
    "donovan": ["donny"],
    "dorothy": ["dodie", "doll", "dolly", "dot"],
    "douglas": ["doug", "dougie"],
    "duncan": ["dunky"],
    "eben": ["ebbie"],
    "ebenezer": ["ebbie"],
    "eberhard": ["ebbe"],
    "ed": ["eddie", "eddy"],
    "edgar": ["ed", "eddie", "eddy"],
    "edward": ["ed", "eddie", "eddy", "ned", "ted", "teddie", "teddy"],
    "edwin": ["ed", "eddie", "eddy"],
    "elbert": ["elbie"],
    "eleanor": ["ellie", "nell", "nellie", "nelly"],
    "elias": ["eli"],
    "elijah": ["eli"],
    "elisabeth": ["beth"],
    "elizabeth": ["bess", "bessie", "bessy", "beth", "betsy", "bette", "betty", "ellie", "izzy", "libby", "lilibet", "liz", "lizzie", "lizzy"],
    "Elzbieta": ["liz"],
    "ellen": ["nell"],
    "emily": ["em", "emmie", "emmy", "millie", "milly"],
    "emma": ["em", "emmie", "emmy"],
    "emmanuel": ["mannie", "manny"],
    "emmy": ["emmie"],
    "enrico": ["rico"],
    "ernest": ["ernie"],
    "esther": ["essie", "hettie", "hetty"],
    "ethel": ["eth"],
    "eugene": ["gene"],
    "euphemia": ["effie", "phemie"],
    "eva": ["evie"],
    "evangeline": ["evie"],
    "eve": ["evie"],
    "evelyn": ["evie", "lyn"],
    "ezekiel": ["zeke"],
    "faith": ["fay"],
    "fay": ["faye"],
    "felicity": ["fliss"],
    "ferdinand": ["ferd", "ferdie", "ferdy", "nandy"],
    "fernando": ["nando"],
    "fiona": ["fi"],
    "flora": ["flo", "florrie"],
    "florence": ["flo", "florrie", "flossie", "flossy", "floy"],
    "frances": ["fannie", "fanny", "fay", "fran", "francie", "frannie", "franny"],
    "francine": ["fran", "francie", "frannie", "franny"],
    "francis": ["fran", "frank", "frankie"],
    "franklin": ["frank", "frankie"],
    "frederick": ["fred", "freddie", "freddy"],
    "gabriel": ["gabby", "gabe"],
    "gabriela": ["gabby"],
    "gabriella": ["gabby"],
    "gabrielle": ["gabby"],
    "gail": ["gayle"],
    "gareth": ["gar", "gare", "gary"],
    "gary": ["gar", "gare"],
    "gavin": ["gav"],
    "geoffrey": ["geoff"],
    "george": ["geordie", "georgie"],
    "georgia": ["george"],
    "georgina": ["george", "georgie", "gina"],
    "gerald": ["gerry", "jerry"],
    "geraldine": ["geri", "gerri", "gerry", "jeri", "jerri", "jerry"],
    "gerard": ["ged", "gerry", "jerry"],
    "gertrude": ["gertie", "trudie", "trudy"],
    "gilbert": ["gil"],
    "gilford": ["gil"],
    "gillian": ["gill"],
    "gina": ["gena"],
    "ginevra": ["ginny"],
    "ginger": ["ging"],
    "gordon": ["don", "gordie", "gordy"],
    "grace": ["gracie"],
    "gregor": ["greg"],
    "gregory": ["greg", "gregg"],
    "greig": ["greg"],
    "giusippe": ["joe"],
    "gustav": ["gus"],
    "hallie": ["hally"],
    "harold": ["hal", "harry"],
    "harriet": ["hallie", "hally", "hattie"],
    "harry": ["hal"],
    "harvey": ["harve"],
    "heathcliff": ["heath"],
    "helen": ["nell", "nellie", "nelly"],
    "henricus": ["harry"],
    "henrietta": ["etta", "hettie", "hetty"],
    "henry": ["hal", "hank", "harry"],
    "herbert": ["herb", "herbie"],
    "hester": ["hettie", "hetty"],
    "hetty": ["hettie"],
    "howard": ["howie"],
    "hugh": ["huey", "hughie"],
    "humphrey": ["huffie"],
    "ienken": ["jill"],
    "ignacio": ["nacho", "nacio"],
    "immanuel": ["mannie", "manny"],
    "indiana": ["indy"],
    "isaac": ["ike"],
    "isabel": ["bel", "bell", "isy", "izzie", "izzy"],
    "isabella": ["bel", "bell", "bella", "isy", "izzie", "izzy"],
    "isabelle": ["bel", "bell", "belle", "isy", "izzie", "izzy"],
    "isidore": ["dore", "izzy"],
    "isobel": ["bel", "bell", "isy", "izzie", "izzy"],
    "ivan": ["van"],
    "jack": ["jackie", "jacko", "jacky", "jak"],
    "jacob": ["jake"],
    "jacqueline": ["jackie", "jacky", "jacqui"],
    "jacquelyn": ["jackie"],
    "james": ["jake", "jamey", "jamie", "jay", "jem", "jim", "jimbo", "jimi", "jimmer", "jimmie", "jimmy"],
    "jane": ["janey", "janie", "jennie", "jenny"],
    "janet": ["jan", "jessie"],
    "janice": ["jan"],
    "jasmine": ["jas", "jaz"],
    "jason": ["jace", "jay"],
    "jean": ["jeanie", "jeannie"],
    "jeannette": ["nettie"],
    "jedidiah": ["jed"],
    "jeffery": ["jeff"],
    "jeffrey": ["jeff"],
    "jemima": ["jem"],
    "jennifer": ["jen", "jenn", "jenni", "jennie", "jenny"],
    "jeremiah": ["jem", "jeremy", "jerry"],
    "jeremy": ["jem", "jerry", "jez"],
    "jerilyn": ["jeri", "jerri", "jerry"],
    "jerrold": ["jerry"],
    "jesse": ["jess"],
    "jessica": ["jess", "jessie"],
    "jessie": ["jesse"],
    "jill": ["jilly"],
    "jillian": ["jilly"],
    "jim": ["jimmy"],
    "jimmy": ["jim", "jimi"],
    "jo": ["jodi", "jodie", "jody"],
    "joan": ["jo"],
    "joanna": ["jo", "jojo"],
    "joanne": ["jojo"],
    "joe": ["jo", "jody", "joey"],
    "johannes": ["joh"],
    "john": ["jackie", "jacky", "johnnie", "johnny", "jon"],
    "jonathan": ["jon", "jonty"],
    "joseph": ["jo", "jody", "joe", "joey", "pepe"],
    "josephine": ["jo", "josie"],
    "joshua": ["josh"],
    "judith": ["jodi", "jodie", "jody", "judie", "judy"],
    "judy": ["judie"],
    "julia": ["jilly", "juley"],
    "juliana": ["julie"],
    "june": ["junie"],
    "kate": ["cate"],
    "katharine": ["kat"],
    "katherine": ["kat", "kate", "kath", "kathi", "kathie", "kathy", "kati", "katy", "kay", "kit", "kitty"],
    "kathryn": ["kat", "kate", "kath", "kathi", "kathie", "kathy", "kati", "katy", "kay", "kit", "kitty"],
    "kathleen": ["kath", "kathy"],
    "kathy": ["kathi", "kathie"],
    "katie": ["caty", "kati"],
    "kay": ["kaye"],
    "kennedy": ["ken", "kenny"],
    "kenneth": ["ken", "kenny"],
    "kerry": ["kez"],
    "kester": ["kit"],
    "kevin": ["kev"],
    "kimberley": ["kimmy"],
    "kimberly": ["kimmy"],
    "kipling": ["kip"],
    "kristina": ["kris", "krissie", "krissy"],
    "kristine": ["krissy"],
    "kurt": ["curt"],
    "kurtis": ["kurt"],
    "lafayette": ["lafe"],
    "laura": ["lauri", "laurie", "lorie"],
    "laurence": ["larry", "lauren", "laurie", "law"],
    "laurie": ["lauri", "lorie"],
    "lavinia": ["vinnie"],
    "lawrence": ["larry", "lauren", "laurie", "law", "lawrie"],
    "leo": ["lee"],
    "leonard": ["len", "lenny"],
    "leopold": ["leo"],
    "leroy": ["lee"],
    "leslie": ["les", "lez"],
    "leslie": ["les"],
    "lester": ["les"],
    "letitia": ["lettie", "letty", "tish", "tisha", "titty"],
    "lilian": ["lil"],
    "lillian": ["lil", "lillie", "lilly"],
    "lily": ["lillie", "lilly"],
    "linda": ["lindie", "lindy"],
    "lindsay": ["lindie", "lindy"],
    "louis": ["lou", "louie"],
    "louisa": ["lou", "lula", "lulu"],
    "louise": ["lou", "lulu"],
    "lucinda": ["cindy"],
    "lucy": ["lu", "lulu"],
    "luke": ["lucky"],
    "luzviminda": ["Luz"],
    "lydia": ["liddy"],
    "lynn": ["lynne"],
    "madeleine": ["maddy"],
    "madonna": ["madge"],
    "magdalene": ["maddy"],
    "maggie": ["mags"],
    "mamie": ["mayme"],
    "marcus": ["marc"],
    "margaret": ["madge", "mae", "maggie", "mags", "maisie", "marge", "margie", "may", "meg", "meggie", "peg", "peggy"],
    "margerita": ["rita"],
    "marilla": ["rilla"],
    "mark": ["marky"],
    "martha": ["marty", "mattie", "matty"],
    "martin": ["marty"],
    "martina": ["marty"],
    "marvin": ["marv"],
    "mary": ["jo", "mae", "mamie", "may", "mayme", "mollie", "polly"],
    "matilda": ["mattie", "matty", "tillie", "tilly"],
    "matthew": ["mat", "matt", "matty"],
    "mattie": ["matty"],
    "maud": ["maudie"],
    "maude": ["maudie"],
    "maurice": ["mo", "moe", "mossie"],
    "mavis": ["mamie"],
    "max": ["mac", "mack", "maxie", "maxy"],
    "maximilian": ["max"],
    "maxine": ["maxie", "maxy"],
    "maxwell": ["max"],
    "megan": ["meg", "meggie"],
    "melanie": ["mel"],
    "melinda": ["mel", "mendy", "mindy"],
    "melissa": ["mel", "missy"],
    "melody": ["mel"],
    "melvin": ["mel"],
    "melvyn": ["mel"],
    "mervin": ["merv"],
    "mervyn": ["merv"],
    "michael": ["mick", "mickey", "micky", "mike", "mikey"],
    "michaela": ["mickey", "micky"],
    "michelle": ["micky", "shell"],
    "mick": ["mickey", "micky"],
    "mikaela": ["mickey", "micky"],
    "mike": ["mickey", "micky", "mikey"],
    "mildred": ["millie", "milly"],
    "millicent": ["millie", "milly"],
    "miranda": ["mindy", "randy"],
    "mitchell": ["mitch"],
    "molly": ["mollie", "polly"],
    "montague": ["monty"],
    "montgomery": ["monty"],
    "morris": ["mo", "moe"],
    "mortimer": ["mort"],
    "morton": ["mort"],
    "moses": ["mo", "moe"],
    "moshe": ["moishy"],
    "murdoch": ["murdy"],
    "nancy": ["nan", "nance"],
    "natalie": ["nat", "talie"],
    "natalya": ["nat"],
    "natasha": ["nat", "tash", "tasha"],
    "nathan": ["nat", "nate"],
    "nathanael": ["nat", "nate"],
    "nathaniel": ["nat", "nate"],
    "nellie": ["nelly"],
    "newton": ["newt"],
    "nichola": ["nic"],
    "nicholas": ["coll", "nic", "nick", "nicky"],
    "nicola": ["nic", "nicki", "nicky", "nikki"],
    "nicole": ["nic", "nicki", "nicky", "nikki"],
    "norlaily": ["laily"],
    "noreen": ["norrie"],
    "norette": ["norrie"],
    "norman": ["norm"],
    "olive": ["ollie"],
    "oliver": ["oli", "ollie"],
    "olivia": ["livvy", "ollie"],
    "olwen": ["ollie"],
    "oswald": ["ozzie", "ozzy"],
    "pamela": ["pam", "pammy"],
    "patricia": ["pat", "patsy", "patti", "patty", "tricia", "trish", "trisha", "trixie"],
    "patrick": ["paddy", "pat", "patsy"],
    "patterson": ["pat"],
    "patty": ["patti"],
    "paul": ["paulie"],
    "peggy": ["peg"],
    "penelope": ["penny"],
    "perceval": ["perce", "percy"],
    "percival": ["perce", "percy"],
    "peregrine": ["perry"],
    "peter": ["pete", "petey"],
    "philip": ["phil", "pip"],
    "philippa": ["phil", "pip", "pippa"],
    "phillip": ["phil", "pip"],
    "philomena": ["phil"],
    "priscilla": ["prissy"],
    "prudence": ["prue"],
    "quincy": ["quin"],
    "quinton": ["quin"],
    "rachel": ["rae", "ray"],
    "rae": ["dannie", "ray"],
    "rajendra": ["raj"],
    "ralph": ["ralphie"],
    "randall": ["randy"],
    "randolph": ["randy"],
    "raphael": ["ralph", "raph"],
    "ray": ["rae"],
    "raymond": ["ray"],
    "rebecca": ["becca", "becki", "becky", "beki"],
    "regina": ["gina"],
    "reginald": ["reg", "reggie"],
    "reuben": ["rube"],
    "ricardo": ["rick", "rico"],
    "richard": ["dick", "dickey", "dickie", "dickon", "dicky", "hick", "rich", "richie", "rick", "rickey", "ricki", "rickie", "ricky"],
    "rickie": ["ricki"],
    "ricky": ["rickey", "ricki"],
    "robbie": ["rabbie"],
    "robert": ["bob", "bobbi", "bobbie", "bobby", "rabbie", "rob", "robbie", "robby"],
    "roberta": ["bobbi", "bobbie", "bobby", "robbi", "robby"],
    "rocco": ["rocky"],
    "roderick": ["rod"],
    "rodney": ["rod"],
    "roger": ["rog"],
    "ronald": ["ron", "ronnie", "ronny"],
    "ronnie": ["ronny"],
    "rose": ["rosie"],
    "rosemary": ["rose", "rosie"],
    "roxana": ["roxy"],
    "roxane": ["roxy"],
    "roxanne": ["roxy"],
    "rudolph": ["rudy"],
    "russell": ["russ"],
    "saffron": ["saffy"],
    "sally": ["sal"],
    "salvador": ["sal"],
    "samantha": ["sam", "sammi", "sammie", "sammy"],
    "samson": ["sam"],
    "samuel": ["sal", "sam", "sammie", "sammy"],
    "sandra": ["sandy"],
    "sarah": ["sadie", "sallie", "sally"],
    "sathamoney": ["sath"],
    "scott": ["scottie", "scotty"],
    "sebastian": ["bastian", "seb"],
    "sharon": ["shaz"],
    "shelley": ["shelly"],
    "shirley": ["shirl"],
    "shlomo": ["shloime"],
    "siddharth": ["sid"],
    "siddhartha": ["sid"],
    "sidney": ["sid"],
    "simeon": ["sim"],
    "simon": ["si", "sim"],
    "solomon": ["sol"],
    "spencer": ["spence"],
    "stacey": ["stace"],
    "stanford": ["stan"],
    "stanislas": ["stan"],
    "stanislav": ["stan"],
    "stanislaw": ["stan"],
    "stanley": ["stan"],
    "stephanie": ["stef", "steff", "steffi", "steffie", "steffy", "steph", "stephie", "steve", "stevie"],
    "stephen": ["steenie", "steve", "stevie"],
    "steven": ["steve", "stevie"],
    "stewart": ["stew", "stewie"],
    "stuart": ["stu"],
    "susan": ["su", "sue", "susie", "susy", "suzi", "suzie", "suzy"],
    "susanna": ["suze"],
    "suzanne": ["suze"],
    "sydney": ["sid", "syd"],
    "sylvester": ["sly"],
    "tallulah": ["lula"],
    "tammy": ["tammie"],
    "teddy": ["teddie"],
    "terence": ["tel", "terry"],
    "teresa": ["terri", "terrie", "terry", "tess"],
    "terry": ["terri", "terrie"],
    "thaddaeus": ["tad", "thad", "thady"],
    "thaddeus": ["tad", "thad", "thady"],
    "theodora": ["theo"],
    "theodore": ["ted", "teddie", "teddy", "theo"],
    "theresa": ["teri", "terry", "tessy"],
    "thomas": ["thom", "tom", "tommy"],
    "timon": ["tim"],
    "timothy": ["tim", "timmy"],
    "tisha": ["tish"],
    "tobiah": ["toby"],
    "tobias": ["toby"],
    "tracy": ["trace"],
    "trevor": ["trev"],
    "trudy": ["trudi", "trudie"],
    "tyler": ["ty"],
    "tyron": ["ty"],
    "tyrone": ["ty"],
    "tyson": ["ty"],
    "valda": ["val"],
    "valentine": ["val"],
    "valerie": ["val"],
    "valeria": ["val"],
    "valerius": ["val"],
    "vance": ["van"],
    "vernon": ["vern"],
    "victor": ["vic", "vick"],
    "victoria": ["vic", "vickey", "vicki", "vickie", "vicky", "vikki"],
    "vincent": ["vin", "vince", "vinnie"],
    "viola": ["vi"],
    "violet": ["vi"],
    "virginia": ["ging", "ginny"],
    "vivian": ["viv"],
    "wallace": ["wally"],
    "walter": ["wally", "walt"],
    "wesley": ["wes"],
    "wilber": ["wil", "will"],
    "wilbur": ["wil", "will"],
    "wilford": ["wil", "will"],
    "wilfred": ["fred", "wil", "wilf", "will"],
    "wilhelmina": ["minnie", "willie"],
    "willard": ["wil", "will"],
    "william": ["bill", "billy", "will", "willie", "willy"],
    "winifred": ["winnie"],
    "winston": ["winnie"],
    "woodrow": ["woody"],
    "zachariah": ["zack", "zak"],
    "zachary": ["zach", "zack", "zak"],
    "zedekiah": ["zed"]
}`

const Nickname_To_Name = `{"abbey": "abigail", "abbi": "abigail", "abbie": "abigail", "abby": "abigail", "abi": "abigail", "gail": "abigail", "gayle": "gail", "addie": "adeline", "aggie": "agnes", "nessie": "agnes", "al": "alfred", "lex": "alexander", "lexi": "alexis", "alec": "alexander", "alex": "alexander", "aleck": "alexander", "elsie": "alexander", "sander": "alexander", "sandy": "sandra", "sawney": "alexander", "xan": "alexander", "zander": "alexander", "allie": "alvin", "ally": "alison", "aly": "alyssa", "lexie": "alexis", "alf": "alfred", "alfie": "alfred", "alfy": "alfred", "fonz": "alfonzo", "fonzie": "alfonzo", "fred": "wilfred", "ali": "alyson", "lonnie": "alonzo", "lonny": "alonzo", "alvie": "alvin", "alvy": "alvin", "mandy": "amanda", "ambers": "ambrose", "ambie": "ambrose", "amby": "ambrose", "brose": "ambrose", "millie": "millicent", "andy": "andrew", "drew": "andrew", "ange": "angelina", "angie": "angelina", "anj": "anjana", "annie": "ann", "jo": "mary", "nan": "nancy", "nannie": "annie", "nanny": "anne", "belle": "isabelle", "nettie": "jeannette", "ant": "anthony", "tony": "anthony", "toni": "antoinette", "tonia": "antonia", "bella": "isabella", "archie": "archibald", "arnie": "arnold", "arfie": "arthur", "art": "arthur", "artie": "arthur", "arty": "arthur", "ash": "ashley", "audie": "audrey", "augie": "augustus", "gus": "gustav", "babs": "barbara", "barb": "barbara", "barbie": "barbara", "barney": "bernard", "baz": "basil", "bazza": "basil", "bart": "bartholomew", "batt": "bartholomew", "bea": "beatrice", "trixie": "patricia", "becki": "rebecca", "beki": "rebecca", "bell": "isobel", "ben": "benjamin", "benji": "benjamin", "benjie": "benjamin", "benjy": "benjamin", "bennie": "bernice", "benny": "bernice", "bernie": "bernard", "bertie": "bertram", "bessy": "elizabeth", "beth": "elizabeth", "bette": "elizabeth", "brad": "bradly", "bri": "brian", "biddie": "bridget", "biddy": "bridget", "cade": "caden", "cal": "caroline", "cale": "caleb", "cam": "cameron", "milly": "millicent", "camp": "campbell", "candi": "candace", "candy": "candice", "callie": "caroline", "caro": "caroline", "carrie": "caroline", "cary": "caroline", "lyn": "evelyn", "cass": "cassandra", "cassie": "cassandra", "cat": "catherine", "cate": "kate", "cath": "catherine", "cathi": "cathy", "cathie": "cathy", "cathy": "catherine", "catie": "catherine", "caty": "katie", "kate": "kathryn", "katie": "catherine", "ceecee": "cecilia", "sissy": "cecilia", "cecie": "cecily", "chad": "charles", "charley": "charlotte", "charlie": "charlotte", "chas": "charles", "chaz": "charles", "chip": "christopher", "chuck": "charles", "lottie": "charlotte", "chance": "chauncey", "cher": "cherilyn", "chet": "chester", "chris": "christopher", "kris": "kristina", "christy": "christina", "kristi": "christina", "kristie": "christina", "kristy": "christina", "chrissie": "christine", "chrissy": "christine", "krissy": "kristine", "christie": "christopher", "kit": "kester", "chuckie": "chuck", "chucky": "chuck", "clay": "clayton", "clem": "clementine", "clemmie": "clement", "cliff": "clifford", "clint": "clinton", "con": "cornelius", "connie": "constance", "corrie": "cornelia", "curt": "kurt", "cindy": "lucinda", "cy": "cyrus", "dan": "danny", "dani": "danny", "danny": "danielle", "danni": "danielle", "dannie": "rae", "dave": "david", "davey": "david", "davie": "david", "davo": "david", "davy": "david", "deb": "debra", "debbie": "debra", "debby": "debra", "denny": "dennis", "den": "dennis", "del": "derek", "jeremy": "jeremiah", "des": "desmond", "di": "diane", "dermot": "diarmuid", "diarmi": "diarmid", "dickie": "richard", "dolly": "dorothy", "dom": "dominic", "donnie": "donald", "donny": "donovan", "don": "gordon", "dodie": "dorothy", "doll": "dorothy", "dot": "dorothy", "doug": "douglas", "dougie": "douglas", "dunky": "duncan", "ebbie": "ebenezer", "ebbe": "eberhard", "eddie": "edwin", "eddy": "edwin", "ed": "edwin", "ned": "edward", "ted": "theodore", "teddie": "theodore", "teddy": "theodore", "elbie": "elbert", "ellie": "elizabeth", "nell": "helen", "nellie": "helen", "nelly": "nellie", "eli": "elijah", "bess": "elizabeth", "bessie": "elizabeth", "betsy": "elizabeth", "betty": "elizabeth", "izzy": "isobel", "libby": "elizabeth", "lilibet": "elizabeth", "liz": "Elzbieta", "lizzie": "elizabeth", "lizzy": "elizabeth", "em": "emma", "emmie": "emmy", "emmy": "emma", "mannie": "immanuel", "manny": "immanuel", "rico": "ricardo", "ernie": "ernest", "essie": "esther", "hettie": "hetty", "hetty": "hester", "eth": "ethel", "gene": "eugene", "effie": "euphemia", "phemie": "euphemia", "evie": "evelyn", "zeke": "ezekiel", "fay": "frances", "faye": "fay", "fliss": "felicity", "ferd": "ferdinand", "ferdie": "ferdinand", "ferdy": "ferdinand", "nandy": "ferdinand", "nando": "fernando", "fi": "fiona", "flo": "florence", "florrie": "florence", "flossie": "florence", "flossy": "florence", "floy": "florence", "fannie": "frances", "fanny": "frances", "fran": "francis", "francie": "francine", "frannie": "francine", "franny": "francine", "frank": "franklin", "frankie": "franklin", "freddie": "frederick", "freddy": "frederick", "gabby": "gabrielle", "gabe": "gabriel", "gar": "gary", "gare": "gary", "gary": "gareth", "gav": "gavin", "geoff": "geoffrey", "geordie": "george", "georgie": "georgina", "george": "georgina", "gina": "regina", "gerry": "gerard", "jerry": "jerrold", "geri": "geraldine", "gerri": "geraldine", "jeri": "jerilyn", "jerri": "jerilyn", "ged": "gerard", "gertie": "gertrude", "trudie": "trudy", "trudy": "gertrude", "gil": "gilford", "gill": "gillian", "gena": "gina", "ginny": "virginia", "ging": "virginia", "gordie": "gordon", "gordy": "gordon", "gracie": "grace", "greg": "greig", "gregg": "gregory", "joe": "joseph", "hally": "harriet", "hal": "henry", "harry": "henry", "hallie": "harriet", "hattie": "harriet", "harve": "harvey", "heath": "heathcliff", "etta": "henrietta", "hank": "henry", "herb": "herbert", "herbie": "herbert", "howie": "howard", "huey": "hugh", "hughie": "hugh", "huffie": "humphrey", "jill": "ienken", "nacho": "ignacio", "nacio": "ignacio", "indy": "indiana", "ike": "isaac", "bel": "isobel", "isy": "isobel", "izzie": "isobel", "dore": "isidore", "van": "vance", "jackie": "john", "jacko": "jack", "jacky": "john", "jak": "jack", "jake": "james", "jacqui": "jacqueline", "jamey": "james", "jamie": "james", "jay": "jason", "jem": "jeremy", "jim": "jimmy", "jimbo": "james", "jimi": "jimmy", "jimmer": "james", "jimmie": "james", "jimmy": "jim", "janey": "jane", "janie": "jane", "jennie": "jennifer", "jenny": "jennifer", "jan": "janice", "jessie": "jessica", "jas": "jasmine", "jaz": "jasmine", "jace": "jason", "jeanie": "jean", "jeannie": "jean", "jed": "jedidiah", "jeff": "jeffrey", "jen": "jennifer", "jenn": "jennifer", "jenni": "jennifer", "jez": "jeremy", "jess": "jessica", "jesse": "jessie", "jilly": "julia", "jodi": "judith", "jodie": "judith", "jody": "judith", "jojo": "joanne", "joey": "joseph", "joh": "johannes", "johnnie": "john", "johnny": "john", "jon": "jonathan", "jonty": "jonathan", "pepe": "joseph", "josie": "josephine", "josh": "joshua", "judie": "judy", "judy": "judith", "juley": "julia", "julie": "juliana", "junie": "june", "kat": "kathryn", "kath": "kathleen", "kathi": "kathy", "kathie": "kathy", "kathy": "kathleen", "kati": "katie", "katy": "kathryn", "kay": "kathryn", "kitty": "kathryn", "kaye": "kay", "ken": "kenneth", "kenny": "kenneth", "kez": "kerry", "kev": "kevin", "kimmy": "kimberly", "kip": "kipling", "krissie": "kristina", "kurt": "kurtis", "lafe": "lafayette", "lauri": "laurie", "laurie": "lawrence", "lorie": "laurie", "larry": "lawrence", "lauren": "lawrence", "law": "lawrence", "vinnie": "vincent", "lawrie": "lawrence", "lee": "leroy", "len": "leonard", "lenny": "leonard", "leo": "leopold", "les": "lester", "lettie": "letitia", "letty": "letitia", "tish": "tisha", "tisha": "letitia", "titty": "letitia", "lil": "lillian", "lillie": "lily", "lilly": "lily", "lindie": "lindsay", "lindy": "lindsay", "lou": "louise", "louie": "louis", "lula": "tallulah", "lulu": "lucy", "lu": "lucy", "lucky": "luke", "Luz": "luzviminda", "liddy": "lydia", "lynne": "lynn", "maddy": "magdalene", "madge": "margaret", "mags": "margaret", "mayme": "mary", "marc": "marcus", "mae": "mary", "maggie": "margaret", "maisie": "margaret", "marge": "margaret", "margie": "margaret", "may": "mary", "meg": "megan", "meggie": "megan", "peg": "peggy", "peggy": "margaret", "rita": "margerita", "rilla": "marilla", "marky": "mark", "marty": "martina", "mattie": "matilda", "matty": "mattie", "marv": "marvin", "mamie": "mavis", "mollie": "molly", "polly": "molly", "tillie": "matilda", "tilly": "matilda", "mat": "matthew", "matt": "matthew", "maudie": "maude", "mo": "moses", "moe": "moses", "mossie": "maurice", "mac": "max", "mack": "max", "maxie": "maxine", "maxy": "maxine", "max": "maxwell", "mel": "melvyn", "mendy": "melinda", "mindy": "miranda", "missy": "melissa", "merv": "mervyn", "mick": "michael", "mickey": "mike", "micky": "mike", "mike": "michael", "mikey": "mike", "shell": "michelle", "randy": "randolph", "mitch": "mitchell", "monty": "montgomery", "mort": "morton", "moishy": "moshe", "murdy": "murdoch", "nance": "nancy", "nat": "nathaniel", "talie": "natalie", "tash": "natasha", "tasha": "natasha", "nate": "nathaniel", "newt": "newton", "nic": "nicole", "coll": "nicholas", "nick": "nicholas", "nicky": "nicole", "nicki": "nicole", "nikki": "nicole", "laily": "norlaily", "norrie": "norette", "norm": "norman", "ollie": "olwen", "oli": "oliver", "livvy": "olivia", "ozzie": "oswald", "ozzy": "oswald", "pam": "pamela", "pammy": "pamela", "pat": "patterson", "patsy": "patrick", "patti": "patty", "patty": "patricia", "tricia": "patricia", "trish": "patricia", "trisha": "patricia", "paddy": "patrick", "paulie": "paul", "penny": "penelope", "perce": "percival", "percy": "percival", "perry": "peregrine", "pete": "peter", "petey": "peter", "phil": "philomena", "pip": "phillip", "pippa": "philippa", "prissy": "priscilla", "prue": "prudence", "quin": "quinton", "rae": "ray", "ray": "raymond", "raj": "rajendra", "ralphie": "ralph", "ralph": "raphael", "raph": "raphael", "becca": "rebecca", "becky": "rebecca", "reg": "reginald", "reggie": "reginald", "rube": "reuben", "rick": "richard", "dick": "richard", "dickey": "richard", "dickon": "richard", "dicky": "richard", "hick": "richard", "rich": "richard", "richie": "richard", "rickey": "ricky", "ricki": "ricky", "rickie": "richard", "ricky": "richard", "rabbie": "robert", "bob": "robert", "bobbi": "roberta", "bobbie": "roberta", "bobby": "roberta", "rob": "robert", "robbie": "robert", "robby": "roberta", "robbi": "roberta", "rocky": "rocco", "rod": "rodney", "rog": "roger", "ron": "ronald", "ronnie": "ronald", "ronny": "ronnie", "rosie": "rosemary", "rose": "rosemary", "roxy": "roxanne", "rudy": "rudolph", "russ": "russell", "saffy": "saffron", "sal": "samuel", "sam": "samuel", "sammi": "samantha", "sammie": "samuel", "sammy": "samuel", "sadie": "sarah", "sallie": "sarah", "sally": "sarah", "sath": "sathamoney", "scottie": "scott", "scotty": "scott", "bastian": "sebastian", "seb": "sebastian", "shaz": "sharon", "shelly": "shelley", "shirl": "shirley", "shloime": "shlomo", "sid": "sydney", "sim": "simon", "si": "simon", "sol": "solomon", "spence": "spencer", "stace": "stacey", "stan": "stanley", "stef": "stephanie", "steff": "stephanie", "steffi": "stephanie", "steffie": "stephanie", "steffy": "stephanie", "steph": "stephanie", "stephie": "stephanie", "steve": "steven", "stevie": "steven", "steenie": "stephen", "stew": "stewart", "stewie": "stewart", "stu": "stuart", "su": "susan", "sue": "susan", "susie": "susan", "susy": "susan", "suzi": "susan", "suzie": "susan", "suzy": "susan", "suze": "suzanne", "syd": "sydney", "sly": "sylvester", "tammie": "tammy", "tel": "terence", "terry": "theresa", "terri": "terry", "terrie": "terry", "tess": "teresa", "tad": "thaddeus", "thad": "thaddeus", "thady": "thaddeus", "theo": "theodore", "teri": "theresa", "tessy": "theresa", "thom": "thomas", "tom": "thomas", "tommy": "thomas", "tim": "timothy", "timmy": "timothy", "toby": "tobias", "trace": "tracy", "trev": "trevor", "trudi": "trudy", "ty": "tyson", "val": "valerius", "vern": "vernon", "vic": "victoria", "vick": "victor", "vickey": "victoria", "vicki": "victoria", "vickie": "victoria", "vicky": "victoria", "vikki": "victoria", "vin": "vincent", "vince": "vincent", "vi": "violet", "viv": "vivian", "wally": "walter", "walt": "walter", "wes": "wesley", "wil": "willard", "will": "william", "wilf": "wilfred", "minnie": "wilhelmina", "willie": "william", "bill": "william", "billy": "william", "willy": "william", "winnie": "winston", "woody": "woodrow", "zack": "zachary", "zak": "zachary", "zach": "zachary", "zed": "zedekiah"}`
