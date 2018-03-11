package edcode

/*
function Decode(Src: string; var Dest: string): Boolean;
var
	StringInfo: TStringInfo;
	sDest: string;
	sSrc: string;
begin
	Result := False;
	Dest := '';
	sDest := ReverseStr(Trim(Src));
	try
		sDest := DecryStrHex(sDest, IntToStr(398432431{240621028}));
	except
		Exit;
	end;
	FillChar(StringInfo, SizeOf(TStringInfo), 0);
	Move(sDest[1], StringInfo, Length(sDest));
	sSrc := StrPas(@StringInfo.sString);
	if (GetUniCode(sSrc) = StringInfo.nUniCode) and (Length(sSrc) = StringInfo.btLength) then begin
		Dest := sSrc;
		Result := True;
	end;
end;
*/

func Decode(src string) (string, error) {
	sDest := ReverseStrBytes(Trim(src))
	sDest, err := DecryStrHex(sDest, COMMON_DECRYPT_KEY)
	if err != nil {
		return "", err;
	}
	return "", nil
}

func DecryStrHex(src, key string) (string, error) {
	//TODO
	return "", nil
}