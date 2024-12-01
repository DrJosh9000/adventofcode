Attribute VB_Name = "mod1B"
Option Base 0

Sub Day1B()
    Open "C:\code\aoc2024\1.txt" For Input As #1
    Dim a, b As Long
    Dim alist() As Long
    Dim blist(100000) As Long
    Dim N As Long
    Do Until EOF(1)
        Input #1, a, b
        ReDim Preserve alist(N + 1)
        alist(N) = a
        blist(b) = blist(b) + 1
        N = N + 1
    Loop
    
    Dim sum As Long
    For i = LBound(alist) To UBound(alist)
        sum = sum + alist(i) * blist(alist(i))
    Next i
    MsgBox sum
End Sub

