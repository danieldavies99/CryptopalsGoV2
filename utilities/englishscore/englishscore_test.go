package englishscore

import (
	"sort"
	"testing"
)

func TestScoreText(t *testing.T) {
	// this test is a bit weird, could probably improve
	// at some point but I don't want to hard code
	// scores in, as it doesn't actually matter what the
	// specific scores are. only that the english scores
	// better than all the others

	testPhrases := [][]byte{
		[]byte("Decoded Ieeacdm*GI-y*fcao*k*zedn*el*hkied"),
		[]byte("Maaeg`i.CM)}.bgek.o.~a{`j.ah.loma`"),
		[]byte("Bnnjhof!LB&r!mhjd!`!qntoe!ng!c`bno"),
		[]byte("Hdd`bel+FH,x+gb`n+j+{d~eo+dm+ijhde"),
		[]byte("Occgebk,AO+,`egi,m,|cybh,cj,nmocb"),
		[]byte("Cooking MC's like a pound of bacon"), // if scoring function works, after sort this should be first
		[]byte("Xttpru|;VX<h;wrp~;z;ktnu;t};yzxtu"),
		[]byte("AmmikleOAqnkigcrmwlfmd`caml"),
		[]byte("^rrvtsz=P^:n=qtvx=|=mrhsy=r{=|~rs"),
		[]byte("_sswur{<Q_;o<puwy<}<lsirx<sz<~}sr"),
	}

	type phraseScore struct {
		phrase []byte
		score  float32
	}

	phraseScores := []phraseScore{}

	for _, phrase := range testPhrases {
		score := ScoreText(phrase)
		phraseScores = append(phraseScores, phraseScore{phrase, score})
	}

	// sort by score
	sort.Slice(phraseScores, func(i, j int) bool {
		return phraseScores[i].score > phraseScores[j].score
	})

	if string(phraseScores[0].phrase) != "Cooking MC's like a pound of bacon" {
		t.Errorf("Expected to find english text, got %s", string(phraseScores[0].phrase))
	}
}
