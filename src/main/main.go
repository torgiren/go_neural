package main
import "neural"
//import "fmt"
import "genetic"
import "sync"

func calcFitness(wg *sync.WaitGroup, genome *genetic.Genome, nets []*neural.Network, i int, inputs [][]int, outputs [][]int) int {
    defer wg.Done()
    var points = 255 * len(inputs) * len(inputs[0])
    var calc []int
    var penalty int
    var diff int
    neural.LoadNetwork(nets[i], genome.GetGenome())
    for j,v := range inputs {
        penalty = 0
        calc = neural.CalcNetwork(nets[i],v)
        for k,v2 := range calc {
            diff = v2 - outputs[j][k]
            if diff < 0 {
                diff *= -1
            }
            penalty += diff
        }
        points -= penalty
//        fmt.Println(inputs[j], calc, outputs[j], penalty, points)
    }
    genome.SetFitness(points)
    return 0
}

func main() {
    var input_num = 2;
    var hidden_num = 4;
    var hidden_size = 6;
    var output_num = 2;
    var pop_size = 200
    var gen_number = 10000
    var nets []*neural.Network
    for i := 0; i< pop_size; i++ {
        nets = append(nets, neural.NewNetwork(input_num, hidden_num, hidden_size, output_num))
    }
//    if neural.LoadNetwork(net, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,1, 1,1,1,1}) != 0 {
//        fmt.Println("ERROR: Load error")
//        return
//    }
//    neural.PrintNet1(net)
//    fmt.Println("dump: ", neural.DumpNetwork(net))
//    fmt.Println("calc: ", neural.CalcNetwork(net, []int{1, 1}))
 //   neural.LoadNetwork(net, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,0, 0,0,0,0})
 //   fmt.Println("calc: ", neural.CalcNetwork(net, []int{1, 1}))
 //   neural.LoadNetwork(net, []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,-1, -1,-1,-1,-1})
 //   fmt.Println("calc: ", neural.CalcNetwork(net, []int{1, 1}))
//    neural.PrintNet1(net)

//    fmt.Println(neural.GetNetworkLen(nets[0]))

    var test_input [][]int = [][]int{ {3,4},{20,4},{8,30}, {30,42}, {200,34}, {255, 2}, {243, 22}, {212, 23}, {30, 250}, {24, 223}, {2, 210}, {9,239} }
    var test_output [][]int = [][]int{ {255,127}, {255, 127}, {255,127},{255,127}, {127,255}, {127,255}, {127,255},{127,255}, {127,255},{127,255}, {127,255}, {127,255} }

    var pop *genetic.Population = genetic.CreatePopulation(pop_size, neural.GetNetworkLen(nets[0]))
    //fmt.Println(pop)
    ////genetic.PrintPopulation(pop)
    //fmt.Println(genetic.GetIndividuals(pop))

    var wg sync.WaitGroup

    for gen := 0; gen<=gen_number; gen++ {
        for i, v := range genetic.GetIndividuals(pop) {
            wg.Add(1)
            go calcFitness(&wg, v, nets, i, test_input, test_output)
        }
        wg.Wait()
        genetic.SortPopulation(pop)
        if gen % 10 == 0 {
            genetic.PrintBest(pop)
        }
        //if gen == 0 || gen == 500{
        //    genetic.PrintPopulation(pop)
        //}
        //genetic.PrintBest(pop)
        //genetic.PrintPopulation(pop)
        pop = genetic.Reproduce(pop)
    }
}
