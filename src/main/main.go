package main
import "neural"
import "fmt"
import "genetic"
import "sync"

func calcFitness(wg *sync.WaitGroup, genome *genetic.Genome, nets []*neural.Network, id int) int {
    defer wg.Done()
    genome.SetFitness(2)
    return 0
}

func main() {
    var input_num = 2;
    var hidden_num = 1;
    var hidden_size = 4;
    var output_num = 2;
    var pop_size = 10
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
    var pop *genetic.Population = genetic.CreatePopulation(pop_size, neural.GetNetworkLen(nets[0]))
    fmt.Println(pop)
    genetic.PrintPopulation(pop)
    fmt.Println(genetic.GetIndividuals(pop))

    var wg sync.WaitGroup

    for i, v := range genetic.GetIndividuals(pop) {
        wg.Add(1)
        go calcFitness(&wg, v, nets, i)
    }
    wg.Wait()
    genetic.PrintPopulation(pop)
}
