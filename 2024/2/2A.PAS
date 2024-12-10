program day2a;

function sign(x: Integer): Integer;
begin
  if x = 0 then sign := 0;
  if x < 0 then sign := -1;
  if x > 0 then sign := 1;
end;

var
   fin: Text;
   ls: Integer;
   a, b: Integer;
   bs: string;
   l: string;
   i, si: Integer;
   safe: Boolean;
   count: Integer;
   code: Integer;

begin
  Assign(fin, 'C:\CODE\AOC\2.txt');
  Reset(fin);
  while not Eof(fin) do begin
    Readln(fin, l);
    a := -1;
    safe := True;
    si := 1;
    ls := 0;
    for i := 1 to Length(l) do begin
      if (l[i] <> ' ') and (i < Length(l)) then Continue;
      if i = Length(l) then
        bs := Copy(l, si, i-si+1)
      else
        bs := Copy(l, si, i-si);
      Val(bs, b, code);
      if code <> 0 then begin
        Writeln('bad number', bs);
        Halt(1)
      end;
      si := i + 1;
      if a = -1 then begin
        a := b;
        Continue
      end;
      if (a = b) or (a-b > 3) or (b-a > 3) then begin
        safe := false;
        Break
      end;
      if ls = 0 then begin
        ls := sign(b-a);
        a := b;
        Continue
      end;
      if ls <> sign(b-a) then begin
        safe := false;
        Break
      end;
      a := b
    end;
    if safe then Inc(count);
  end;
  Writeln(count);
  Readln
end.