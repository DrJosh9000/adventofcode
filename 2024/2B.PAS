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
   ints: array[1..20] of Integer;
   intsl: Integer;
   skip: Integer;
   i, si: Integer;
   safe, anysafe: Boolean;
   count: Integer;
   code: Integer;

begin
  Assign(fin, 'C:\CODE\AOC\2.txt');
  Reset(fin);
  while not Eof(fin) do begin
    Readln(fin, l);
    intsl := 0;
    si := 1;
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
      Inc(intsl);
      ints[intsl] := b;
    end;

    anysafe := False;
    for skip := 1 to intsl do begin
      a := -1;
      safe := True;
      ls := 0;
      for i := 1 to intsl do begin
        if i = skip then Continue;
        b := ints[i];
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
      if safe then begin
        anysafe := True;
        Break
      end
    end;
    if anysafe then Inc(count)
  end;
  Writeln(count);
  Readln
end.