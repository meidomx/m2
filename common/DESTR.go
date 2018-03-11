package edcode

import (
	"errors"
)

var (
	ERROR_NOT_A_HEX_STRING = errors.New("not a hex string")
)

func HexToInt(src string) (uint32, error) {
	/*
	  var
    I, Res: Integer;
    Ch: Char;
  begin
    Res := 0;
    for I := 0 to Length(Hex) - 1 do begin
      Ch := Hex[I + 1];
      if (Ch >= '0') and (Ch <= '9') then
        Res := Res * 16 + Ord(Ch) - Ord('0')
      else if (Ch >= 'A') and (Ch <= 'F') then
        Res := Res * 16 + Ord(Ch) - Ord('A') + 10
      else if (Ch >= 'a') and (Ch <= 'f') then
        Res := Res * 16 + Ord(Ch) - Ord('a') + 10
      else raise Exception.Create('Error: not a Hex String');
    end;
    Result := Res;
  end;
	 */
	b := []byte(src)
	r := uint32(0)
	for _, b := range b {
		if b >= '0' && b <= '9' {
			r = r * 16 +
		}
	}
}