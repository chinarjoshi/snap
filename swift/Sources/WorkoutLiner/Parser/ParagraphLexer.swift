// Generated from ./Paragraph.g4 by ANTLR 4.13.2
import Antlr4

open class ParagraphLexer: Lexer {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = ParagraphLexer._ATN.getNumberOfDecisions()
          for i in 0..<length {
          	    decisionToDFA.append(DFA(ParagraphLexer._ATN.getDecisionState(i)!, i))
          }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	static let NUMBER=1, X=2, BY=3, WORD=4, NEWLINE=5, WS=6

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
		"NUMBER", "X", "BY", "WORD", "NEWLINE", "WS"
	]

	private static let _LITERAL_NAMES: [String?] = [
		nil, nil, nil, nil, nil, "'\\n'"
	]
	private static let _SYMBOLIC_NAMES: [String?] = [
		nil, "NUMBER", "X", "BY", "WORD", "NEWLINE", "WS"
	]
	public
	static let VOCABULARY = Vocabulary(_LITERAL_NAMES, _SYMBOLIC_NAMES)


	override open
	func getVocabulary() -> Vocabulary {
		return ParagraphLexer.VOCABULARY
	}

	public
	required init(_ input: CharStream) {
	    RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION)
		super.init(input)
		_interp = LexerATNSimulator(self, ParagraphLexer._ATN, ParagraphLexer._decisionToDFA, ParagraphLexer._sharedContextCache)
	}

	override open
	func getGrammarFileName() -> String { return "Paragraph.g4" }

	override open
	func getRuleNames() -> [String] { return ParagraphLexer.ruleNames }

	override open
	func getSerializedATN() -> [Int] { return ParagraphLexer._serializedATN }

	override open
	func getChannelNames() -> [String] { return ParagraphLexer.channelNames }

	override open
	func getModeNames() -> [String] { return ParagraphLexer.modeNames }

	override open
	func getATN() -> ATN { return ParagraphLexer._ATN }

	static let _serializedATN:[Int] = [
		4,0,6,37,6,-1,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,1,0,4,0,
		15,8,0,11,0,12,0,16,1,1,1,1,1,2,1,2,1,2,1,3,4,3,25,8,3,11,3,12,3,26,1,
		4,1,4,1,5,4,5,32,8,5,11,5,12,5,33,1,5,1,5,0,0,6,1,1,3,2,5,3,7,4,9,5,11,
		6,1,0,6,1,0,48,57,2,0,88,88,120,120,2,0,66,66,98,98,2,0,89,89,121,121,
		2,0,65,90,97,122,2,0,9,9,32,32,39,0,1,1,0,0,0,0,3,1,0,0,0,0,5,1,0,0,0,
		0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,1,14,1,0,0,0,3,18,1,0,0,0,5,20,1,
		0,0,0,7,24,1,0,0,0,9,28,1,0,0,0,11,31,1,0,0,0,13,15,7,0,0,0,14,13,1,0,
		0,0,15,16,1,0,0,0,16,14,1,0,0,0,16,17,1,0,0,0,17,2,1,0,0,0,18,19,7,1,0,
		0,19,4,1,0,0,0,20,21,7,2,0,0,21,22,7,3,0,0,22,6,1,0,0,0,23,25,7,4,0,0,
		24,23,1,0,0,0,25,26,1,0,0,0,26,24,1,0,0,0,26,27,1,0,0,0,27,8,1,0,0,0,28,
		29,5,10,0,0,29,10,1,0,0,0,30,32,7,5,0,0,31,30,1,0,0,0,32,33,1,0,0,0,33,
		31,1,0,0,0,33,34,1,0,0,0,34,35,1,0,0,0,35,36,6,5,0,0,36,12,1,0,0,0,4,0,
		16,26,33,1,6,0,0
	]

	public
	static let _ATN: ATN = try! ATNDeserializer().deserialize(_serializedATN)
}