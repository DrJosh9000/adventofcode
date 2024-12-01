Attribute VB_Name = "mod1a"
Option Base 0

Sub Day1A()
    Open "C:\code\aoc2024\1.txt" For Input As #1
    Dim a, b As Long
    Dim alist() As Variant, blist() As Variant
    Dim N As Long
    Do Until EOF(1)
        Input #1, a, b
        ReDim Preserve alist(N + 1)
        ReDim Preserve blist(N + 1)
        alist(N) = a
        blist(N) = b
        N = N + 1
    Loop

    Call VBCore.SortArray(alist)
    Call VBCore.SortArray(blist)
    
    Dim sum As Long
    For i = LBound(alist) To UBound(alist)
        If alist(i) < blist(i) Then
            sum = sum + blist(i) - alist(i)
        Else
            sum = sum + alist(i) - blist(i)
        End If
    Next i
    MsgBox sum
End Sub
