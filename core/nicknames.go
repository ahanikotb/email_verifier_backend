package core

import (
	"encoding/json"
)

func FindNickNames(name string) (bool, []string) {
	result := []string{}
	names := map[string][]string{}

	json.Unmarshal([]byte(JSON_DATA), &names)
	value, exists := names[name]
	if exists {
		return exists, value
	}

	return false, result
}

const JSON_DATA = `{
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