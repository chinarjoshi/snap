// Generated from Workoutliner.g4 by ANTLR 4.13.2
import Antlr4

open class WorkoutlinerParser: Parser {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = WorkoutlinerParser._ATN.getNumberOfDecisions()
          for i in 0..<length {
            decisionToDFA.append(DFA(WorkoutlinerParser._ATN.getDecisionState(i)!, i))
           }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	enum Tokens: Int {
		case EOF = -1, NUMBER = 1, X = 2, BY = 3, STAR = 4, WORD = 5, NEWLINE = 6, 
                 WS = 7
	}

	public
	static let RULE_paragraph = 0, RULE_line = 1, RULE_continuationLine = 2, 
            RULE_supersetLine = 3, RULE_supersetName = 4, RULE_exerciseLine = 5, 
            RULE_proseLine = 6, RULE_numericToken = 7, RULE_token = 8, RULE_byExpr = 9, 
            RULE_multiplier = 10, RULE_number = 11, RULE_note = 12

	public
	static let ruleNames: [String] = [
		"paragraph", "line", "continuationLine", "supersetLine", "supersetName", 
		"exerciseLine", "proseLine", "numericToken", "token", "byExpr", "multiplier", 
		"number", "note"
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
	func getGrammarFileName() -> String { return "Workoutliner.g4" }

	override open
	func getRuleNames() -> [String] { return WorkoutlinerParser.ruleNames }

	override open
	func getSerializedATN() -> [Int] { return WorkoutlinerParser._serializedATN }

	override open
	func getATN() -> ATN { return WorkoutlinerParser._ATN }


	override open
	func getVocabulary() -> Vocabulary {
	    return WorkoutlinerParser.VOCABULARY
	}

	override public
	init(_ input:TokenStream) throws {
	    RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION)
		try super.init(input)
		_interp = ParserATNSimulator(self,WorkoutlinerParser._ATN,WorkoutlinerParser._decisionToDFA, WorkoutlinerParser._sharedContextCache)
	}


	public class ParagraphContext: ParserRuleContext {
			open
			func EOF() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.EOF.rawValue, 0)
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
			return WorkoutlinerParser.RULE_paragraph
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterParagraph(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitParagraph(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitParagraph(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 0, WorkoutlinerParser.RULE_paragraph)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(27) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(26)
		 		try line()


		 		setState(29); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (_la == WorkoutlinerParser.Tokens.NUMBER.rawValue || _la == WorkoutlinerParser.Tokens.WORD.rawValue)
		 	setState(31)
		 	try match(WorkoutlinerParser.Tokens.EOF.rawValue)

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
			return WorkoutlinerParser.RULE_line
		}
	}
	public class ExerciseLineAltContext: LineContext {
			open
			func exerciseLine() -> ExerciseLineContext? {
				return getRuleContext(ExerciseLineContext.self, 0)
			}
			open
			func NEWLINE() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterExerciseLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitExerciseLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitExerciseLineAlt(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitExerciseLineAlt(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	public class ContinuationLineAltContext: LineContext {
			open
			func continuationLine() -> ContinuationLineContext? {
				return getRuleContext(ContinuationLineContext.self, 0)
			}
			open
			func NEWLINE() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterContinuationLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitContinuationLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitContinuationLineAlt(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitContinuationLineAlt(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	public class SupersetLineAltContext: LineContext {
			open
			func supersetLine() -> SupersetLineContext? {
				return getRuleContext(SupersetLineContext.self, 0)
			}
			open
			func NEWLINE() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterSupersetLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitSupersetLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitSupersetLineAlt(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitSupersetLineAlt(self)
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
				return getToken(WorkoutlinerParser.Tokens.NEWLINE.rawValue, 0)
			}

		public
		init(_ ctx: LineContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterProseLineAlt(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitProseLineAlt(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitProseLineAlt(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 2, WorkoutlinerParser.RULE_line)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(49)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,5, _ctx)) {
		 	case 1:
		 		_localctx =  SupersetLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(33)
		 		try supersetLine()
		 		setState(35)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == WorkoutlinerParser.Tokens.NEWLINE.rawValue) {
		 			setState(34)
		 			try match(WorkoutlinerParser.Tokens.NEWLINE.rawValue)

		 		}


		 		break
		 	case 2:
		 		_localctx =  ExerciseLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(37)
		 		try exerciseLine()
		 		setState(39)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == WorkoutlinerParser.Tokens.NEWLINE.rawValue) {
		 			setState(38)
		 			try match(WorkoutlinerParser.Tokens.NEWLINE.rawValue)

		 		}


		 		break
		 	case 3:
		 		_localctx =  ContinuationLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 3)
		 		setState(41)
		 		try continuationLine()
		 		setState(43)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == WorkoutlinerParser.Tokens.NEWLINE.rawValue) {
		 			setState(42)
		 			try match(WorkoutlinerParser.Tokens.NEWLINE.rawValue)

		 		}


		 		break
		 	case 4:
		 		_localctx =  ProseLineAltContext(_localctx);
		 		try enterOuterAlt(_localctx, 4)
		 		setState(45)
		 		try proseLine()
		 		setState(47)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (_la == WorkoutlinerParser.Tokens.NEWLINE.rawValue) {
		 			setState(46)
		 			try match(WorkoutlinerParser.Tokens.NEWLINE.rawValue)

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

	public class ContinuationLineContext: ParserRuleContext {
			open
			func numericToken() -> NumericTokenContext? {
				return getRuleContext(NumericTokenContext.self, 0)
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
			return WorkoutlinerParser.RULE_continuationLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterContinuationLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitContinuationLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitContinuationLine(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitContinuationLine(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func continuationLine() throws -> ContinuationLineContext {
		var _localctx: ContinuationLineContext
		_localctx = ContinuationLineContext(_ctx, getState())
		try enterRule(_localctx, 4, WorkoutlinerParser.RULE_continuationLine)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(51)
		 	try numericToken()
		 	setState(55)
		 	try _errHandler.sync(self)
		 	_alt = try getInterpreter().adaptivePredict(_input,6,_ctx)
		 	while (_alt != 2 && _alt != ATN.INVALID_ALT_NUMBER) {
		 		if ( _alt==1 ) {
		 			setState(52)
		 			try token()

		 	 
		 		}
		 		setState(57)
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,6,_ctx)
		 	}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class SupersetLineContext: ParserRuleContext {
			open
			func supersetName() -> [SupersetNameContext] {
				return getRuleContexts(SupersetNameContext.self)
			}
			open
			func supersetName(_ i: Int) -> SupersetNameContext? {
				return getRuleContext(SupersetNameContext.self, i)
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
			return WorkoutlinerParser.RULE_supersetLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterSupersetLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitSupersetLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitSupersetLine(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitSupersetLine(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func supersetLine() throws -> SupersetLineContext {
		var _localctx: SupersetLineContext
		_localctx = SupersetLineContext(_ctx, getState())
		try enterRule(_localctx, 6, WorkoutlinerParser.RULE_supersetLine)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(58)
		 	try supersetName()
		 	setState(59)
		 	try supersetName()
		 	setState(61); 
		 	try _errHandler.sync(self)
		 	_alt = 1;
		 	repeat {
		 		switch (_alt) {
		 		case 1:
		 			setState(60)
		 			try token()


		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(63); 
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,7,_ctx)
		 	} while (_alt != 2 && _alt !=  ATN.INVALID_ALT_NUMBER)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class SupersetNameContext: ParserRuleContext {
			open
			func STAR() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.STAR.rawValue, 0)
			}
			open
			func WORD() -> [TerminalNode] {
				return getTokens(WorkoutlinerParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.WORD.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return WorkoutlinerParser.RULE_supersetName
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterSupersetName(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitSupersetName(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitSupersetName(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
			    return visitor.visitSupersetName(self)
			}
			else {
			     return visitor.visitChildren(self)
			}
		}
	}
	@discardableResult
	 open func supersetName() throws -> SupersetNameContext {
		var _localctx: SupersetNameContext
		_localctx = SupersetNameContext(_ctx, getState())
		try enterRule(_localctx, 8, WorkoutlinerParser.RULE_supersetName)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(66) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(65)
		 		try match(WorkoutlinerParser.Tokens.WORD.rawValue)


		 		setState(68); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (_la == WorkoutlinerParser.Tokens.WORD.rawValue)
		 	setState(70)
		 	try match(WorkoutlinerParser.Tokens.STAR.rawValue)

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
				return getTokens(WorkoutlinerParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.WORD.rawValue, i)
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
			return WorkoutlinerParser.RULE_exerciseLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterExerciseLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitExerciseLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitExerciseLine(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 10, WorkoutlinerParser.RULE_exerciseLine)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(73) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(72)
		 		try match(WorkoutlinerParser.Tokens.WORD.rawValue)


		 		setState(75); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (_la == WorkoutlinerParser.Tokens.WORD.rawValue)
		 	setState(77)
		 	try numericToken()
		 	setState(81)
		 	try _errHandler.sync(self)
		 	_alt = try getInterpreter().adaptivePredict(_input,10,_ctx)
		 	while (_alt != 2 && _alt != ATN.INVALID_ALT_NUMBER) {
		 		if ( _alt==1 ) {
		 			setState(78)
		 			try token()

		 	 
		 		}
		 		setState(83)
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,10,_ctx)
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
				return getTokens(WorkoutlinerParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.WORD.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return WorkoutlinerParser.RULE_proseLine
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterProseLine(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitProseLine(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitProseLine(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 12, WorkoutlinerParser.RULE_proseLine)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(85); 
		 	try _errHandler.sync(self)
		 	_alt = 1;
		 	repeat {
		 		switch (_alt) {
		 		case 1:
		 			setState(84)
		 			try match(WorkoutlinerParser.Tokens.WORD.rawValue)


		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(87); 
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
			return WorkoutlinerParser.RULE_numericToken
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterNumericToken(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitNumericToken(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitNumericToken(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 14, WorkoutlinerParser.RULE_numericToken)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(92)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,12, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(89)
		 		try byExpr()

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(90)
		 		try multiplier()

		 		break
		 	case 3:
		 		try enterOuterAlt(_localctx, 3)
		 		setState(91)
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
			return WorkoutlinerParser.RULE_token
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterToken(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitToken(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitToken(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 16, WorkoutlinerParser.RULE_token)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(96)
		 	try _errHandler.sync(self)
		 	switch (WorkoutlinerParser.Tokens(rawValue: try _input.LA(1))!) {
		 	case .NUMBER:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(94)
		 		try numericToken()

		 		break

		 	case .WORD:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(95)
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
			return WorkoutlinerParser.RULE_byExpr
		}
	}
	public class TwoPartByContext: ByExprContext {
		public var reps: Token!
		public var weight: Token!
			open
			func BY() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.BY.rawValue, 0)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(WorkoutlinerParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: ByExprContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterTwoPartBy(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitTwoPartBy(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitTwoPartBy(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
				return getTokens(WorkoutlinerParser.Tokens.BY.rawValue)
			}
			open
			func BY(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.BY.rawValue, i)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(WorkoutlinerParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: ByExprContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterThreePartBy(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitThreePartBy(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitThreePartBy(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 18, WorkoutlinerParser.RULE_byExpr)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(106)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,14, _ctx)) {
		 	case 1:
		 		_localctx =  TwoPartByContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(98)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(TwoPartByContext.self).reps = assignmentValue
		 		     }()

		 		setState(99)
		 		try match(WorkoutlinerParser.Tokens.BY.rawValue)
		 		setState(100)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(TwoPartByContext.self).weight = assignmentValue
		 		     }()


		 		break
		 	case 2:
		 		_localctx =  ThreePartByContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(101)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(ThreePartByContext.self).sets = assignmentValue
		 		     }()

		 		setState(102)
		 		try match(WorkoutlinerParser.Tokens.BY.rawValue)
		 		setState(103)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(ThreePartByContext.self).reps = assignmentValue
		 		     }()

		 		setState(104)
		 		try match(WorkoutlinerParser.Tokens.BY.rawValue)
		 		setState(105)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
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
			return WorkoutlinerParser.RULE_multiplier
		}
	}
	public class PartialMultiplierContext: MultiplierContext {
		public var sets: Token!
		public var reps: Token!
			open
			func X() -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.X.rawValue, 0)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(WorkoutlinerParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: MultiplierContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterPartialMultiplier(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitPartialMultiplier(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitPartialMultiplier(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
				return getTokens(WorkoutlinerParser.Tokens.X.rawValue)
			}
			open
			func X(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.X.rawValue, i)
			}
			open
			func NUMBER() -> [TerminalNode] {
				return getTokens(WorkoutlinerParser.Tokens.NUMBER.rawValue)
			}
			open
			func NUMBER(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.NUMBER.rawValue, i)
			}

		public
		init(_ ctx: MultiplierContext) {
			super.init()
			copyFrom(ctx)
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterFullMultiplier(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitFullMultiplier(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitFullMultiplier(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 20, WorkoutlinerParser.RULE_multiplier)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(116)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,15, _ctx)) {
		 	case 1:
		 		_localctx =  PartialMultiplierContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(108)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(PartialMultiplierContext.self).sets = assignmentValue
		 		     }()

		 		setState(109)
		 		try match(WorkoutlinerParser.Tokens.X.rawValue)
		 		setState(110)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(PartialMultiplierContext.self).reps = assignmentValue
		 		     }()


		 		break
		 	case 2:
		 		_localctx =  FullMultiplierContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(111)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(FullMultiplierContext.self).sets = assignmentValue
		 		     }()

		 		setState(112)
		 		try match(WorkoutlinerParser.Tokens.X.rawValue)
		 		setState(113)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
		 				_localctx.castdown(FullMultiplierContext.self).reps = assignmentValue
		 		     }()

		 		setState(114)
		 		try match(WorkoutlinerParser.Tokens.X.rawValue)
		 		setState(115)
		 		try {
		 				let assignmentValue = try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)
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
				return getToken(WorkoutlinerParser.Tokens.NUMBER.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return WorkoutlinerParser.RULE_number
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterNumber(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitNumber(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitNumber(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 22, WorkoutlinerParser.RULE_number)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(118)
		 	try match(WorkoutlinerParser.Tokens.NUMBER.rawValue)

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
				return getTokens(WorkoutlinerParser.Tokens.WORD.rawValue)
			}
			open
			func WORD(_ i:Int) -> TerminalNode? {
				return getToken(WorkoutlinerParser.Tokens.WORD.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return WorkoutlinerParser.RULE_note
		}
		override open
		func enterRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.enterNote(self)
			}
		}
		override open
		func exitRule(_ listener: ParseTreeListener) {
			if let listener = listener as? WorkoutlinerListener {
				listener.exitNote(self)
			}
		}
		override open
		func accept<T>(_ visitor: ParseTreeVisitor<T>) -> T? {
			if let visitor = visitor as? WorkoutlinerVisitor {
			    return visitor.visitNote(self)
			}
			else if let visitor = visitor as? WorkoutlinerBaseVisitor {
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
		try enterRule(_localctx, 24, WorkoutlinerParser.RULE_note)
		defer {
	    		try! exitRule()
	    }
		do {
			var _alt:Int
		 	try enterOuterAlt(_localctx, 1)
		 	setState(121); 
		 	try _errHandler.sync(self)
		 	_alt = 1;
		 	repeat {
		 		switch (_alt) {
		 		case 1:
		 			setState(120)
		 			try match(WorkoutlinerParser.Tokens.WORD.rawValue)


		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(123); 
		 		try _errHandler.sync(self)
		 		_alt = try getInterpreter().adaptivePredict(_input,16,_ctx)
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
		4,1,7,126,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,
		7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,1,0,4,0,28,8,0,11,0,
		12,0,29,1,0,1,0,1,1,1,1,3,1,36,8,1,1,1,1,1,3,1,40,8,1,1,1,1,1,3,1,44,8,
		1,1,1,1,1,3,1,48,8,1,3,1,50,8,1,1,2,1,2,5,2,54,8,2,10,2,12,2,57,9,2,1,
		3,1,3,1,3,4,3,62,8,3,11,3,12,3,63,1,4,4,4,67,8,4,11,4,12,4,68,1,4,1,4,
		1,5,4,5,74,8,5,11,5,12,5,75,1,5,1,5,5,5,80,8,5,10,5,12,5,83,9,5,1,6,4,
		6,86,8,6,11,6,12,6,87,1,7,1,7,1,7,3,7,93,8,7,1,8,1,8,3,8,97,8,8,1,9,1,
		9,1,9,1,9,1,9,1,9,1,9,1,9,3,9,107,8,9,1,10,1,10,1,10,1,10,1,10,1,10,1,
		10,1,10,3,10,117,8,10,1,11,1,11,1,12,4,12,122,8,12,11,12,12,12,123,1,12,
		0,0,13,0,2,4,6,8,10,12,14,16,18,20,22,24,0,0,132,0,27,1,0,0,0,2,49,1,0,
		0,0,4,51,1,0,0,0,6,58,1,0,0,0,8,66,1,0,0,0,10,73,1,0,0,0,12,85,1,0,0,0,
		14,92,1,0,0,0,16,96,1,0,0,0,18,106,1,0,0,0,20,116,1,0,0,0,22,118,1,0,0,
		0,24,121,1,0,0,0,26,28,3,2,1,0,27,26,1,0,0,0,28,29,1,0,0,0,29,27,1,0,0,
		0,29,30,1,0,0,0,30,31,1,0,0,0,31,32,5,0,0,1,32,1,1,0,0,0,33,35,3,6,3,0,
		34,36,5,6,0,0,35,34,1,0,0,0,35,36,1,0,0,0,36,50,1,0,0,0,37,39,3,10,5,0,
		38,40,5,6,0,0,39,38,1,0,0,0,39,40,1,0,0,0,40,50,1,0,0,0,41,43,3,4,2,0,
		42,44,5,6,0,0,43,42,1,0,0,0,43,44,1,0,0,0,44,50,1,0,0,0,45,47,3,12,6,0,
		46,48,5,6,0,0,47,46,1,0,0,0,47,48,1,0,0,0,48,50,1,0,0,0,49,33,1,0,0,0,
		49,37,1,0,0,0,49,41,1,0,0,0,49,45,1,0,0,0,50,3,1,0,0,0,51,55,3,14,7,0,
		52,54,3,16,8,0,53,52,1,0,0,0,54,57,1,0,0,0,55,53,1,0,0,0,55,56,1,0,0,0,
		56,5,1,0,0,0,57,55,1,0,0,0,58,59,3,8,4,0,59,61,3,8,4,0,60,62,3,16,8,0,
		61,60,1,0,0,0,62,63,1,0,0,0,63,61,1,0,0,0,63,64,1,0,0,0,64,7,1,0,0,0,65,
		67,5,5,0,0,66,65,1,0,0,0,67,68,1,0,0,0,68,66,1,0,0,0,68,69,1,0,0,0,69,
		70,1,0,0,0,70,71,5,4,0,0,71,9,1,0,0,0,72,74,5,5,0,0,73,72,1,0,0,0,74,75,
		1,0,0,0,75,73,1,0,0,0,75,76,1,0,0,0,76,77,1,0,0,0,77,81,3,14,7,0,78,80,
		3,16,8,0,79,78,1,0,0,0,80,83,1,0,0,0,81,79,1,0,0,0,81,82,1,0,0,0,82,11,
		1,0,0,0,83,81,1,0,0,0,84,86,5,5,0,0,85,84,1,0,0,0,86,87,1,0,0,0,87,85,
		1,0,0,0,87,88,1,0,0,0,88,13,1,0,0,0,89,93,3,18,9,0,90,93,3,20,10,0,91,
		93,3,22,11,0,92,89,1,0,0,0,92,90,1,0,0,0,92,91,1,0,0,0,93,15,1,0,0,0,94,
		97,3,14,7,0,95,97,3,24,12,0,96,94,1,0,0,0,96,95,1,0,0,0,97,17,1,0,0,0,
		98,99,5,1,0,0,99,100,5,3,0,0,100,107,5,1,0,0,101,102,5,1,0,0,102,103,5,
		3,0,0,103,104,5,1,0,0,104,105,5,3,0,0,105,107,5,1,0,0,106,98,1,0,0,0,106,
		101,1,0,0,0,107,19,1,0,0,0,108,109,5,1,0,0,109,110,5,2,0,0,110,117,5,1,
		0,0,111,112,5,1,0,0,112,113,5,2,0,0,113,114,5,1,0,0,114,115,5,2,0,0,115,
		117,5,1,0,0,116,108,1,0,0,0,116,111,1,0,0,0,117,21,1,0,0,0,118,119,5,1,
		0,0,119,23,1,0,0,0,120,122,5,5,0,0,121,120,1,0,0,0,122,123,1,0,0,0,123,
		121,1,0,0,0,123,124,1,0,0,0,124,25,1,0,0,0,17,29,35,39,43,47,49,55,63,
		68,75,81,87,92,96,106,116,123
	]

	public
	static let _ATN = try! ATNDeserializer().deserialize(_serializedATN)
}