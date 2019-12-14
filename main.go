package main
import "fmt"
//import "strconv"

type neuron struct {
    bias int
    synapses []*neuron
    values []int
    input bool
    name string
}
type layer struct {
    neurons []* neuron
}

type network struct {
    inputs *layer
    hiddens []*layer
    outputs *layer
    inputs_num int
    hiddnes_num int
    hiddnes_size int
    outputs_num int
}

//func newNeuron(bias int, neurons []*neuron, values []int, name string) *neuron {
func newNeuron(neurons []*neuron) *neuron {
    var a neuron
    a.synapses = neurons
    a.values = make([]int, len(neurons))
    return &a
}
func newInputNeuron() *neuron {
    var a neuron
    a.input = true
    return &a
}

func newLayer(size int) *layer {
    var l layer
    return &l
}

func main() {
    var input_num = 3;
    var hidden_num = 3;
    var hidden_size = 5;
    var output_num = 2;
    var inputs []*neuron;
    var outputs []*neuron;
    var hiddens [][]*neuron;
    for i := 0; i<input_num; i++ {
        inputs = append(inputs, newInputNeuron())
    }
    hiddens = make([][]*neuron, hidden_num)
    for i := 0; i<hidden_num; i++ {
        for j := 0; j<hidden_size; j++ {
            var prev []*neuron
            if i == 0 {
                prev = inputs
            } else {
                prev = hiddens[i-1]
            }
//            var val []int = make([]int, len(prev))
            hiddens[i] = append(hiddens[i], newNeuron(prev))
            //hiddens[i] = append(hiddens[i], newNeuron(10*i+j, prev, val, "neuron_"+strconv.Itoa(i)+"_"+strconv.Itoa(j)))
        }
    }
    for i := 0; i<output_num; i++ {
//        var val []int = make([]int, hidden_size)
        outputs = append(outputs, newNeuron(hiddens[hidden_num-1]))
        //outputs = append(outputs, newNeuron(i, hiddens[hidden_num-1], val, "neuron_"+strconv.Itoa(i)))
    }
    fmt.Println(inputs)
    fmt.Println(inputs[0])
    fmt.Println(hiddens)
    for i,v := range hiddens {
        for j,v2 := range v {
//            fmt.Println(i, j, *v2, &(*v2))
            fmt.Printf("%d %d %p ", i, j, v2)
            fmt.Println(*v2)
        }
    }
    fmt.Println(outputs)
    for j,v2 := range outputs {
        fmt.Printf("%d %p ", j, v2)
        fmt.Println(*v2)
    }
}
