package main
import "fmt"
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

type network struct {
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

func newNetwork(inputs_num, hiddens_num, hiddens_size, outputs_num int) *network {
    var n network
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

func printNet1(net *network) {
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

func dumpNetwork(net *network) []int {
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

func loadNetwork(net *network, data []int) int {
    fmt.Println("data len: ", len(data))
    var net_len int = net.inputs_num * net.hiddens_size + net.hiddens_size +
                      net.hiddens_size * net.hiddens_size * (net.hiddens_num - 1) + net.hiddens_size * (net.hiddens_num - 1) +
                      net.outputs_num * net.hiddens_size + net.outputs_num
    fmt.Println("net len:", net_len)
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


    fmt.Println("AAAA")
    var pos int = 0
    for i := 0; i<net.hiddens_num; i++ {
        var prev *layer
        if i == 0 {
            prev = net.inputs
        } else {
            prev = net.hiddens[i-1]
        }
        fmt.Println(prev)
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
    fmt.Println("BBBB")



    return 0
}

func main() {
    var input_num = 3;
    var hidden_num = 3;
    var hidden_size = 5;
    var output_num = 2;
    var net *network = newNetwork(input_num, hidden_num, hidden_size, output_num)
    if loadNetwork(net, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91}) != 0 {
        fmt.Println("ERROR: Load error")
        return
    }
    printNet1(net)
    fmt.Println("dump: ", dumpNetwork(net))
}
