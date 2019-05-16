package tests

import (
	"testing"
	"tflac_cw/kernel"
	"tflac_cw/token"

	"github.com/stretchr/testify/assert"
)

type TestKernel struct {
	Core *kernel.Kernel
}

func (test *TestKernel) New() {
	tururu := kernel.Kernel{}
	test.Core = tururu.New()
}

func TestInputOneIdentificator(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("test")

	assert.Equal(t, "<type> test;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputTwoIdentificatorWithComma(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("test, test2")

	assert.Equal(t, "<type> test, test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 7, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputTwoIdentificatorWithoutComma(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("test test2")

	assert.Equal(t, "<type> test, test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 7, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithTwoTypesWithCommaWithoutSecondIdentifier(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float test1, int")

	assert.Equal(t, "float test1;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithTwoTypesWithCommaWithSecondIdentifier(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float test1, int test2")

	assert.Equal(t, "float test1,  test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 8, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithPointer1(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float *test1, *test2;")

	assert.Equal(t, "float *test1, *test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 9, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithPointer2(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float *test1, **test2;")

	assert.Equal(t, "float *test1, *test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 9, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithPointer3(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float **test1, ***test2, ***test3;")

	assert.Equal(t, "float **test1, **test2, **test3;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 16, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithPointer4(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float test1, **test2, *test3;")

	assert.Equal(t, "float test1, test2, test3;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 10, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputSentenceWithTwoIdentifierWithComma(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("test1, test2")

	assert.Equal(t, "<type> test1, test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 7, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestInputManyTypes(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("float test1\n int test2;")

	assert.Equal(t, "float test1;\n int test2;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 10, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestRepairTypeFloat(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("flot test1")

	assert.Equal(t, "float test1;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestRepairTypeFloat2(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("flt test1")

	assert.Equal(t, "float test1;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestRepairTypeInt(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("it test1")

	assert.Equal(t, "int test1;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestRepairTypeInt2(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("i test1")

	assert.Equal(t, "<type> test1;", test.Core.RepairActual(), "they should be equaled!")
	assert.Equal(t, 4, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "amount objects in array should be equal")
}

func TestRepairEndStatement(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("flot test1")

	assert.Equal(t, token.FLOAT, test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired[0].Type, "they should be equaled")
	assert.Equal(t, token.ENDSTATEMENT, test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired[len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired)-1].Type, "they should be equaled!")
}

func TestRepairWithErrorSyms(t *testing.T) {
	test := TestKernel{}
	test.New()
	test.Core.Input("flot &#! test1!@#")

	assert.Equal(t, "float  test1;", test.Core.RepairActual(), "they should be equaled")
	assert.Equal(t, 6, len(test.Core.SyntaxAnalys.SyntaxAnalyser.TokensRepaired), "they should be equaled!")
}
