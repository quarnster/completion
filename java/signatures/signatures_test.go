package signatures

import (
	"testing"
)

func TestSignatures(t *testing.T) {
	tests := []string{
		"<E:Ljava/lang/Object;>Ljava/util/AbstractList<TE;>;Ljava/util/List<TE;>;Ljava/util/RandomAccess;Ljava/lang/Cloneable;Ljava/io/Serializable;",
		"(Ljava/util/Collection<+TE;>;)V",
		"()Ljava/util/Enumeration<TE;>;",
		"(I)TE;",
		"()TE;",
		"()TE;",
		"(TE;I)V",
		"(TE;I)V",
		"(TE;)V",
		"<T:Ljava/lang/Object;>([TT;)[TT;",
		"(I)TE;",
		"(I)TE;",
		"(ITE;)TE;",
		"(TE;)Z",
		"(ITE;)V",
		"(I)TE;",
		"(Ljava/util/Collection<*>;)Z",
		"(Ljava/util/Collection<+TE;>;)Z",
		"(Ljava/util/Collection<*>;)Z",
		"(Ljava/util/Collection<*>;)Z",
		"(ILjava/util/Collection<+TE;>;)Z",
		"(II)Ljava/util/List<TE;>;",
		"(I)Ljava/util/ListIterator<TE;>;",
		"()Ljava/util/ListIterator<TE;>;",
		"()Ljava/util/Iterator<TE;>;",
		"Ljava/lang/Object;Ljava/io/Serializable;Ljava/lang/Comparable<Ljava/lang/String;>;Ljava/lang/CharSequence;",
		"Ljava/util/Comparator<Ljava/lang/String;>;"}
	for i := range tests {
		var p SIGNATURES
		t.Log(tests[i], p.Parse(tests[i]))
		t.Log(p.RootNode().String())
		if p.RootNode().Range.End() != len(tests[i]) {
			t.Error(p.Error())
		}
	}
}
