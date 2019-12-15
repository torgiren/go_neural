package neural
import "fmt"
import "math"
//import "strconv"

type neuron struct {
    bias int
    synapses []*neuron
    values []int
    input bool
}
type layer struct {
    neurons []* neuron
}

type Network struct {
    inputs *layer
    hiddens []*layer
    outputs *layer
    inputs_num int
    hiddens_num int
    hiddens_size int
    outputs_num int
}

//func newNeuron(bias int, neurons []*neuron, values []int, name string) *neuron {
func newNeuron(neurons []*neuron) *neuron {
    var a neuron
    a.synapses = neurons
    a.values = make([]int, len(neurons))
    return &a
}

func newLayer(size int, parent *layer) *layer {
    var l layer
    var parent_neurons []*neuron = nil
    if parent != nil {
        parent_neurons = parent.neurons
    }
    l.neurons = make([]*neuron, size)
    for i:= 0; i<size; i++ {

        l.neurons[i] = newNeuron(parent_neurons)
    }
    return &l
}

func NewNetwork(inputs_num, hiddens_num, hiddens_size, outputs_num int) *Network {
    var n Network
    n.inputs_num = inputs_num
    n.hiddens_num = hiddens_num
    n.hiddens_size = hiddens_size
    n.outputs_num = outputs_num
    n.inputs = newLayer(n.inputs_num, nil)
    for _, v := range n.inputs.neurons {
        v.input = true
    }
    n.hiddens = make([]*layer, n.hiddens_num)
    for i := 0; i<n.hiddens_num; i++ {
        var prev *layer
        if i == 0 {
            prev = n.inputs
        } else {
            prev = n.hiddens[i-1]
        }
        n.hiddens[i] = newLayer(n.hiddens_size, prev)
    }
    n.outputs = newLayer(n.outputs_num, n.hiddens[n.hiddens_num-1])
    return &n
}

func PrintNet1(net *Network) {
    fmt.Println(net)
    fmt.Println("inputs", net.inputs)
    for i,v := range net.inputs.neurons {
        fmt.Printf("%d %p ", i, v)
        fmt.Println(*v)
    }
    fmt.Println("hiddens", net.hiddens)
    for i,v := range net.hiddens {
        for j,v2 := range v.neurons {
            fmt.Printf("%d %d %p ", i, j, v2)
            fmt.Println(*v2)
        }
    }
    fmt.Println("outputs", net.outputs)
    for i,v := range net.outputs.neurons {
        fmt.Printf("%d %p ", i, v)
        fmt.Println(*v)
    }
}

func DumpNetwork(net *Network) []int {
    var dump []int
    for _, v := range net.hiddens {
        for _, v2 := range v.neurons {
            dump = append(dump, v2.values...)
            dump = append(dump, v2.bias)
        }
    }
    for _, v := range net.outputs.neurons {
        dump = append(dump, v.values...)
        dump = append(dump, v.bias)
    }
    return dump
}

func LoadNetwork(net *Network, data []int) int {
    var net_len int = net.inputs_num * net.hiddens_size + net.hiddens_size +
                      net.hiddens_size * net.hiddens_size * (net.hiddens_num - 1) + net.hiddens_size * (net.hiddens_num - 1) +
                      net.outputs_num * net.hiddens_size + net.outputs_num
    if len(data) != net_len {
        fmt.Println("ERROR: data len differ from network len: data_len ", len(data), ", net_len ", net_len)
        return -1
    }
//    for i := 0; i < net.hiddens_size; i++ {
//        net.hiddens[0].neurons[i].values = data[i*(net.inputs_num+1):(i+1) * (net.inputs_num+1)- 1]
//        net.hiddens[0].neurons[i].bias = data[(i+1)*(net.inputs_num+1)-1]
//    }
//    var hidden_start int = net.hiddens_size * (net.inputs_num+1)
//    fmt.Println(hidden_start)


    var pos int = 0
    for i := 0; i<net.hiddens_num; i++ {
        var prev *layer
        if i == 0 {
            prev = net.inputs
        } else {
            prev = net.hiddens[i-1]
        }
        for _, v2 := range net.hiddens[i].neurons {
            v2.values = data[pos:pos+len(prev.neurons)]
            v2.bias = data[pos+len(prev.neurons)]
            pos+=len(prev.neurons)+1
        }
    }
    for _, v := range net.outputs.neurons {
        v.values = data[pos:pos+net.hiddens_size]
        v.bias = data[pos+net.hiddens_size]
        pos+=net.hiddens_size+1
    }

    return 0
}

func CalcNetwork(net *Network, inputs []int) []int {
    fmt.Println("calc")
    if len(inputs) != net.inputs_num {
        fmt.Println("ERROR: input values len differ from network inputs num: inputs_values len ", len(inputs), ", inputs_num ", net.inputs_num)
        return []int{-1}
    }
    for i, v := range net.inputs.neurons {
        v.bias = inputs[i]
    }
    var result []int = make([]int, net.outputs_num)
    for i := 0; i<net.outputs_num; i++ {
        result[i] = calcNeuron(net.outputs.neurons[i])
    }
    return result
}

func calcNeuron(neuron *neuron) int {
    var result int = neuron.bias
    if neuron.input {
        return result
    }
    for i := 0; i<len(neuron.synapses); i++ {
        result += neuron.values[i] * calcNeuron(neuron.synapses[i])
    }
    return sigmoid(result)
}
func sigmoid(val int) int {
    var result int = int(1.0 / (1 + math.Exp(-1.0*(float64(val)/5))) * 255)
    return result
}

