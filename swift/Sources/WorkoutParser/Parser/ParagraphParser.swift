// Generated from ./Paragraph.g4 by ANTLR 4.13.2
import Antlr4

open class ParagraphParser: Parser {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = ParagraphParser._ATN.getNumberOfDecisions()
          for i in 0..<length {
            decisionToDFA.append(DFA(ParagraphParser._ATN.getDecisionState(i)!, i))
           }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	enum Tokens: Int {
		case EOF = -1, NUMBER = 1, X = 2, BY = 3, WORD = 4, NEWLINE = 5, WS = 6
	}

	public
	static let RULE_paragraph = 0, RULE_line = 1, RULE_exerciseLine = 2, RULE_proseLine = 3, 
            RULE_numericToken = 4, RULE_token = 5, RULE_byExpr = 6, RULE_multiplier = 7, 
            RULE_number = 8, RULE_note = 9

	public
	static let ruleNames: [String] = [
		"paragraph", "line", "exerciseLine", "proseLine", "numericToken", "token", 
		"byExpr", "multiplier", "number", "note"
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
	func getGrammarFileName() -> String { return "Paragraph.g4" }

	override open
	func getRuleNames() -> [String] { return ParagraphParser.ruleNames }

	override open
	func getSerializedATN() -> [Int] { return ParagraphParser._serializedATN }

	override open
	func getATN() -> ATN { return ParagraphParser._ATN }


	override open
	func getVocabulary() -> Vocabulary {
	    return ParagraphParser.VOCABULARY
	}

	override public
	init(_ input:TokenStream) throws {
	    RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION)
		try super.init(input)
		_interp = ParserATNSimulator(self,ParagraphParser._ATN,ParagraphParser._decisionToDFA, ParagraphParser._sharedContextCache)
	}


	public class ParagraphContext: ParserRuleContext {
			open
			func EOF() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.EOF.rawValue, 0)
			}
			open
			func line() -> [LineContext] {
				return getRuleContexts(LineContext.self)
			}
			open
			func line(_ i: Int) -> LineContext? {
				return getRuleContext(LineContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_paragraph
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterParagraph(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitParagraph(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitParagraph(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitParagraph(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func paragraph() throws -> ParagraphContext {
		var _localctx: ParagraphContext
		_localctx = ParagraphContext(_ctx, getState())
		try enterRule(_localctx, 0, ParagraphParser.RULE_paragraph)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(21) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(20)
		 		try line()


		 		setState(23); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (_la == ParagraphParser.Tokens.WORD.rawValue)
		 	setState(25)
		 	try match(ParagraphParser.Tokens.EOF.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class LineContext: ParserRuleContext {
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_line
		}
	}
	public class ExerciseLineAltContext: LineContext {
			open
			func exerciseLine() -> ExerciseLineContext? {
				return getRuleContext(ExerciseLineContext.self, 0)
			}
			open
			func NEWLINE() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterExerciseLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitExerciseLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitExerciseLineAlt(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitExerciseLineAlt(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	public class ProseLineAltContext: LineContext {
			open
			func proseLine() -> ProseLineContext? {
				return getRuleContext(ProseLineContext.self, 0)
			}
			open
			func NEWLINE() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterProseLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitProseLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitProseLineAlt(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitProseLineAlt(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func line() throws -> LineContext {
		var _localctx: LineContext
		_localctx = LineContext(_ctx, getState())
		try enterRule(_localctx, 2, ParagraphParser.RULE_line)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(35)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,3, _ctx)) {
		 	case 1:
		 		_localctx =  ExerciseLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(27)
		 		try exerciseLine()
		 		setState(29)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == ParagraphParser.Tokens.NEWLINE.rawValue) {
		 			setState(28)
		 			try match(ParagraphParser.Tokens.NEWLINE.rawValue)

		 		}


		 		break
		 	case 2:
		 		_localctx =  ProseLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(31)
		 		try proseLine()
		 		setState(33)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == ParagraphParser.Tokens.NEWLINE.rawValue) {
		 			setState(32)
		 			try match(ParagraphParser.Tokens.NEWLINE.rawValue)

		 		}


		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class ExerciseLineContext: ParserRuleContext {
			open
			func numericToken() -> NumericTokenContext? {
				return getRuleContext(NumericTokenContext.self, 0)
			}
			open
			func WORD() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.WORD.rawValue, i)
			}
			open
			func token() -> [TokenContext] {
				return getRuleContexts(TokenContext.self)
			}
			open
			func token(_ i: Int) -> TokenContext? {
				return getRuleContext(TokenContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_exerciseLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterExerciseLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitExerciseLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitExerciseLine(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitExerciseLine(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func exerciseLine() throws -> ExerciseLineContext {
		var _localctx: ExerciseLineContext
		_localctx = ExerciseLineContext(_ctx, getState())
		try enterRule(_localctx, 4, ParagraphParser.RULE_exerciseLine)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(38) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(37)
		 		try match(ParagraphParser.Tokens.WORD.rawValue)


		 		setState(40); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (_la == ParagraphParser.Tokens.WORD.rawValue)
		 	setState(42)
		 	try numericToken()
		 	setState(46)
		 	try _errHandler.sync(self)
		 	_alt = try getInterpreter().adaptivePredict(_input,5,_ctx)
		 	while (_alt != 2 && _alt != ATN.INVALID_ALT_NUMBER) {
		 		if ( _alt==1 ) {
		 			setState(43)
		 			try token()

		 	 
		 		}
		 		setState(48)
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,5,_ctx)
		 	}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class ProseLineContext: ParserRuleContext {
			open
			func WORD() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.WORD.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_proseLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterProseLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitProseLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitProseLine(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitProseLine(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func proseLine() throws -> ProseLineContext {
		var _localctx: ProseLineContext
		_localctx = ProseLineContext(_ctx, getState())
		try enterRule(_localctx, 6, ParagraphParser.RULE_proseLine)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(50); 
		 	try _errHandler.sync(self)
		 	_alt = 1;
		 	repeat {
		 		switch (_alt) {
		 		case 1:
		 			setState(49)
		 			try match(ParagraphParser.Tokens.WORD.rawValue)


		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(52); 
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,6,_ctx)
		 	} while (_alt != 2 && _alt !=  ATN.INVALID_ALT_NUMBER)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class NumericTokenContext: ParserRuleContext {
			open
			func byExpr() -> ByExprContext? {
				return getRuleContext(ByExprContext.self, 0)
			}
			open
			func multiplier() -> MultiplierContext? {
				return getRuleContext(MultiplierContext.self, 0)
			}
			open
			func number() -> NumberContext? {
				return getRuleContext(NumberContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_numericToken
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterNumericToken(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitNumericToken(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitNumericToken(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitNumericToken(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func numericToken() throws -> NumericTokenContext {
		var _localctx: NumericTokenContext
		_localctx = NumericTokenContext(_ctx, getState())
		try enterRule(_localctx, 8, ParagraphParser.RULE_numericToken)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(57)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,7, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(54)
		 		try byExpr()

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(55)
		 		try multiplier()

		 		break
		 	case 3:
		 		try enterOuterAlt(_localctx, 3)
		 		setState(56)
		 		try number()

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class TokenContext: ParserRuleContext {
			open
			func numericToken() -> NumericTokenContext? {
				return getRuleContext(NumericTokenContext.self, 0)
			}
			open
			func note() -> NoteContext? {
				return getRuleContext(NoteContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_token
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterToken(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitToken(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitToken(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitToken(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func token() throws -> TokenContext {
		var _localctx: TokenContext
		_localctx = TokenContext(_ctx, getState())
		try enterRule(_localctx, 10, ParagraphParser.RULE_token)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(61)
		 	try _errHandler.sync(self)
		 	switch (ParagraphParser.Tokens(rawValue: try _input.LA(1))!) {
		 	case .NUMBER:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(59)
		 		try numericToken()

		 		break

		 	case .WORD:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(60)
		 		try note()

		 		break
		 	default:
		 		throw ANTLRException.recognition(e: NoViableAltException(self))
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class ByExprContext: ParserRuleContext {
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_byExpr
		}
	}
	public class TwoPartByContext: ByExprContext {
		public var reps: Token!
		public var weight: Token!
			open
			func BY() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.BY.rawValue, 0)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: ByExprContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterTwoPartBy(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitTwoPartBy(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitTwoPartBy(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitTwoPartBy(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	public class ThreePartByContext: ByExprContext {
		public var sets: Token!
		public var reps: Token!
		public var weight: Token!
			open
			func BY() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.BY.rawValue)
			}
			open
			func BY(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.BY.rawValue, i)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: ByExprContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterThreePartBy(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitThreePartBy(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitThreePartBy(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitThreePartBy(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func byExpr() throws -> ByExprContext {
		var _localctx: ByExprContext
		_localctx = ByExprContext(_ctx, getState())
		try enterRule(_localctx, 12, ParagraphParser.RULE_byExpr)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(71)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,9, _ctx)) {
		 	case 1:
		 		_localctx =  TwoPartByContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(63)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(TwoPartByContext.self).reps = assignmentValue
		 		     }()

		 		setState(64)
		 		try match(ParagraphParser.Tokens.BY.rawValue)
		 		setState(65)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(TwoPartByContext.self).weight = assignmentValue
		 		     }()


		 		break
		 	case 2:
		 		_localctx =  ThreePartByContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(66)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(ThreePartByContext.self).sets = assignmentValue
		 		     }()

		 		setState(67)
		 		try match(ParagraphParser.Tokens.BY.rawValue)
		 		setState(68)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(ThreePartByContext.self).reps = assignmentValue
		 		     }()

		 		setState(69)
		 		try match(ParagraphParser.Tokens.BY.rawValue)
		 		setState(70)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(ThreePartByContext.self).weight = assignmentValue
		 		     }()


		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class MultiplierContext: ParserRuleContext {
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_multiplier
		}
	}
	public class PartialMultiplierContext: MultiplierContext {
		public var sets: Token!
		public var reps: Token!
			open
			func X() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.X.rawValue, 0)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: MultiplierContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterPartialMultiplier(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitPartialMultiplier(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitPartialMultiplier(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitPartialMultiplier(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	public class FullMultiplierContext: MultiplierContext {
		public var sets: Token!
		public var reps: Token!
		public var weight: Token!
			open
			func X() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.X.rawValue)
			}
			open
			func X(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.X.rawValue, i)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: MultiplierContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterFullMultiplier(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitFullMultiplier(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitFullMultiplier(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitFullMultiplier(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func multiplier() throws -> MultiplierContext {
		var _localctx: MultiplierContext
		_localctx = MultiplierContext(_ctx, getState())
		try enterRule(_localctx, 14, ParagraphParser.RULE_multiplier)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(81)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,10, _ctx)) {
		 	case 1:
		 		_localctx =  PartialMultiplierContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(73)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(PartialMultiplierContext.self).sets = assignmentValue
		 		     }()

		 		setState(74)
		 		try match(ParagraphParser.Tokens.X.rawValue)
		 		setState(75)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(PartialMultiplierContext.self).reps = assignmentValue
		 		     }()


		 		break
		 	case 2:
		 		_localctx =  FullMultiplierContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(76)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(FullMultiplierContext.self).sets = assignmentValue
		 		     }()

		 		setState(77)
		 		try match(ParagraphParser.Tokens.X.rawValue)
		 		setState(78)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(FullMultiplierContext.self).reps = assignmentValue
		 		     }()

		 		setState(79)
		 		try match(ParagraphParser.Tokens.X.rawValue)
		 		setState(80)
		 		try {
		 				let assignmentValue = try match(ParagraphParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(FullMultiplierContext.self).weight = assignmentValue
		 		     }()


		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class NumberContext: ParserRuleContext {
			open
			func NUMBER() -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.NUMBER.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_number
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterNumber(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitNumber(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitNumber(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitNumber(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func number() throws -> NumberContext {
		var _localctx: NumberContext
		_localctx = NumberContext(_ctx, getState())
		try enterRule(_localctx, 16, ParagraphParser.RULE_number)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(83)
		 	try match(ParagraphParser.Tokens.NUMBER.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class NoteContext: ParserRuleContext {
			open
			func WORD() -> [TerminalNode] {
				return getTokens(ParagraphParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(ParagraphParser.Tokens.WORD.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return ParagraphParser.RULE_note
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.enterNote(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? ParagraphListener {
				listener.exitNote(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? ParagraphVisitor {
			    return visitor.visitNote(self)
			}
			else if let visitor = visitor as? ParagraphBaseVisitor {
			    return visitor.visitNote(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func note() throws -> NoteContext {
		var _localctx: NoteContext
		_localctx = NoteContext(_ctx, getState())
		try enterRule(_localctx, 18, ParagraphParser.RULE_note)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(86); 
		 	try _errHandler.sync(self)
		 	_alt = 1;
		 	repeat {
		 		switch (_alt) {
		 		case 1:
		 			setState(85)
		 			try match(ParagraphParser.Tokens.WORD.rawValue)


		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(88); 
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,11,_ctx)
		 	} while (_alt != 2 && _alt !=  ATN.INVALID_ALT_NUMBER)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	static let _serializedATN:[Int] = [
		4,1,6,91,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,
		7,2,8,7,8,2,9,7,9,1,0,4,0,22,8,0,11,0,12,0,23,1,0,1,0,1,1,1,1,3,1,30,8,
		1,1,1,1,1,3,1,34,8,1,3,1,36,8,1,1,2,4,2,39,8,2,11,2,12,2,40,1,2,1,2,5,
		2,45,8,2,10,2,12,2,48,9,2,1,3,4,3,51,8,3,11,3,12,3,52,1,4,1,4,1,4,3,4,
		58,8,4,1,5,1,5,3,5,62,8,5,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,3,6,72,8,6,1,
		7,1,7,1,7,1,7,1,7,1,7,1,7,1,7,3,7,82,8,7,1,8,1,8,1,9,4,9,87,8,9,11,9,12,
		9,88,1,9,0,0,10,0,2,4,6,8,10,12,14,16,18,0,0,93,0,21,1,0,0,0,2,35,1,0,
		0,0,4,38,1,0,0,0,6,50,1,0,0,0,8,57,1,0,0,0,10,61,1,0,0,0,12,71,1,0,0,0,
		14,81,1,0,0,0,16,83,1,0,0,0,18,86,1,0,0,0,20,22,3,2,1,0,21,20,1,0,0,0,
		22,23,1,0,0,0,23,21,1,0,0,0,23,24,1,0,0,0,24,25,1,0,0,0,25,26,5,0,0,1,
		26,1,1,0,0,0,27,29,3,4,2,0,28,30,5,5,0,0,29,28,1,0,0,0,29,30,1,0,0,0,30,
		36,1,0,0,0,31,33,3,6,3,0,32,34,5,5,0,0,33,32,1,0,0,0,33,34,1,0,0,0,34,
		36,1,0,0,0,35,27,1,0,0,0,35,31,1,0,0,0,36,3,1,0,0,0,37,39,5,4,0,0,38,37,
		1,0,0,0,39,40,1,0,0,0,40,38,1,0,0,0,40,41,1,0,0,0,41,42,1,0,0,0,42,46,
		3,8,4,0,43,45,3,10,5,0,44,43,1,0,0,0,45,48,1,0,0,0,46,44,1,0,0,0,46,47,
		1,0,0,0,47,5,1,0,0,0,48,46,1,0,0,0,49,51,5,4,0,0,50,49,1,0,0,0,51,52,1,
		0,0,0,52,50,1,0,0,0,52,53,1,0,0,0,53,7,1,0,0,0,54,58,3,12,6,0,55,58,3,
		14,7,0,56,58,3,16,8,0,57,54,1,0,0,0,57,55,1,0,0,0,57,56,1,0,0,0,58,9,1,
		0,0,0,59,62,3,8,4,0,60,62,3,18,9,0,61,59,1,0,0,0,61,60,1,0,0,0,62,11,1,
		0,0,0,63,64,5,1,0,0,64,65,5,3,0,0,65,72,5,1,0,0,66,67,5,1,0,0,67,68,5,
		3,0,0,68,69,5,1,0,0,69,70,5,3,0,0,70,72,5,1,0,0,71,63,1,0,0,0,71,66,1,
		0,0,0,72,13,1,0,0,0,73,74,5,1,0,0,74,75,5,2,0,0,75,82,5,1,0,0,76,77,5,
		1,0,0,77,78,5,2,0,0,78,79,5,1,0,0,79,80,5,2,0,0,80,82,5,1,0,0,81,73,1,
		0,0,0,81,76,1,0,0,0,82,15,1,0,0,0,83,84,5,1,0,0,84,17,1,0,0,0,85,87,5,
		4,0,0,86,85,1,0,0,0,87,88,1,0,0,0,88,86,1,0,0,0,88,89,1,0,0,0,89,19,1,
		0,0,0,12,23,29,33,35,40,46,52,57,61,71,81,88
	]

	public
	static let _ATN = try! ATNDeserializer().deserialize(_serializedATN)
}