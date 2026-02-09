// Generated from Workoutliner.g4 by ANTLR 4.13.2
import Antlr4

open class WorkoutlinerLexer: Lexer {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = WorkoutlinerLexer._ATN.getNumberOfDecisions()
          for i in 0..<length {
          	    decisionToDFA.append(DFA(WorkoutlinerLexer._ATN.getDecisionState(i)!, i))
          }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	static let NUMBER=1, X=2, BY=3, STAR=4, WORD=5, NEWLINE=6, WS=7

	public
	static let channelNames: [String] = [
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	]

	public
	static let modeNames: [String] = [
		"DEFAULT_MODE"
	]

	public
	static let ruleNames: [String] = [
		"NUMBER", "X", "BY", "STAR", "WORD", "NEWLINE", "WS"
	]

	private static let _LITERAL_NAMES: [String?] = [
		nil, nil, nil, nil, "'*'", nil, "'\\n'"
	]
	private static let _SYMBOLIC_NAMES: [String?] = [
		nil, "NUMBER", "X", "BY", "STAR", "WORD", "NEWLINE", "WS"
	]
	public
	static let VOCABULARY = Vocabulary(_LITERAL_NAMES, _SYMBOLIC_NAMES)


	override open
	func getVocabulary() -> Vocabulary {
		return WorkoutlinerLexer.VOCABULARY
	}

	public
	required init(_ input: CharStream) {
	    RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION)
		super.init(input)
		_interp = LexerATNSimulator(self, WorkoutlinerLexer._ATN, WorkoutlinerLexer._decisionToDFA, WorkoutlinerLexer._sharedContextCache)
	}

	override open
	func getGrammarFileName() -> String { return "Workoutliner.g4" }

	override open
	func getRuleNames() -> [String] { return WorkoutlinerLexer.ruleNames }

	override open
	func getSerializedATN() -> [Int] { return WorkoutlinerLexer._serializedATN }

	override open
	func getChannelNames() -> [String] { return WorkoutlinerLexer.channelNames }

	override open
	func getModeNames() -> [String] { return WorkoutlinerLexer.modeNames }

	override open
	func getATN() -> ATN { return WorkoutlinerLexer._ATN }

	static let _serializedATN:[Int] = [
		4,0,7,41,6,-1,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,
		1,0,4,0,17,8,0,11,0,12,0,18,1,1,1,1,1,2,1,2,1,2,1,3,1,3,1,4,4,4,29,8,4,
		11,4,12,4,30,1,5,1,5,1,6,4,6,36,8,6,11,6,12,6,37,1,6,1,6,0,0,7,1,1,3,2,
		5,3,7,4,9,5,11,6,13,7,1,0,6,1,0,48,57,2,0,88,88,120,120,2,0,66,66,98,98,
		2,0,89,89,121,121,2,0,65,90,97,122,2,0,9,9,32,32,43,0,1,1,0,0,0,0,3,1,
		0,0,0,0,5,1,0,0,0,0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,13,1,0,0,0,1,
		16,1,0,0,0,3,20,1,0,0,0,5,22,1,0,0,0,7,25,1,0,0,0,9,28,1,0,0,0,11,32,1,
		0,0,0,13,35,1,0,0,0,15,17,7,0,0,0,16,15,1,0,0,0,17,18,1,0,0,0,18,16,1,
		0,0,0,18,19,1,0,0,0,19,2,1,0,0,0,20,21,7,1,0,0,21,4,1,0,0,0,22,23,7,2,
		0,0,23,24,7,3,0,0,24,6,1,0,0,0,25,26,5,42,0,0,26,8,1,0,0,0,27,29,7,4,0,
		0,28,27,1,0,0,0,29,30,1,0,0,0,30,28,1,0,0,0,30,31,1,0,0,0,31,10,1,0,0,
		0,32,33,5,10,0,0,33,12,1,0,0,0,34,36,7,5,0,0,35,34,1,0,0,0,36,37,1,0,0,
		0,37,35,1,0,0,0,37,38,1,0,0,0,38,39,1,0,0,0,39,40,6,6,0,0,40,14,1,0,0,
		0,4,0,18,30,37,1,6,0,0
	]

	public
	static let _ATN: ATN = try! ATNDeserializer().deserialize(_serializedATN)
}