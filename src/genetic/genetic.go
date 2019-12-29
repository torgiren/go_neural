package genetic
import "math/rand"
import "time"
import "fmt"
import "sort"

type Genome struct {
    genome_len int
    genome []int
    fitness int

}
type Population struct {
    size int
    individuals []*Genome
}

func createGenome(genome_len int) *Genome{
    var g Genome
    g.genome_len = genome_len
    for i := 0; i< g.genome_len; i++ {
        g.genome = append(g.genome, rand.Intn(510)- 255)
    }
    return &g
}
func createGenomeFromArray(genome []int) *Genome {
    var g Genome
    g.genome_len = len(genome)
    g.genome = genome
    return &g
}

func CreatePopulation(size int, genome_len int) *Population {
    rand.Seed(time.Now().UnixNano())
    var p Population
    p.size = size
    for i:= 0; i<size; i++ {
        p.individuals = append(p.individuals, createGenome(genome_len))
    }
    return &p
}

func PrintPopulation(pop *Population) {
    fmt.Println("Pop size: ", pop.size)
    for i := 0; i< pop.size; i++ {
        PrintGenome(pop.individuals[i])
    }
}

func PrintBest(pop *Population) {
    PrintGenome(pop.individuals[pop.size-1])
}

func PrintGenome(genome *Genome) {
    fmt.Println(genome.genome, genome.fitness)
}

func GetIndividuals(pop *Population) []*Genome{
    return pop.individuals
}

func (genome *Genome) SetFitness(value int) {
    genome.fitness = value
}

func (genome *Genome) GetGenome() []int {
    return genome.genome
}

func SortPopulation(pop *Population) {
    sort.Slice(pop.individuals, func(i int,j int) bool { return pop.individuals[i].fitness < pop.individuals[j].fitness})
}

func selectCandidate(pop* Population) *Genome{
    var sumFitness int = 0
    var tmpFitness int = 0;
    for _, v := range(pop.individuals) {
        //fmt.Println(v.fitness)
        sumFitness += v.fitness
    }
    var randFitness int = rand.Intn(sumFitness)
    //fmt.Println(sumFitness)
    //fmt.Println(randFitness)
    for _, v := range(pop.individuals) {
        tmpFitness += v.fitness
        if randFitness <= tmpFitness {
            return v
        }
    }
    return nil
}

func Reproduce(pop *Population) *Population{
    //fmt.Println("Reproduce")
    var new_pop Population
    for i := 0; i<(pop.size/2) - 1; i++ {
        var cand1 *Genome = selectCandidate(pop)
        var cand2 *Genome = selectCandidate(pop)
        var children [2]*Genome = cross(cand1, cand2)
        //fmt.Println(cand1)
        //fmt.Println(cand2)
        //fmt.Println(children[0])
        //fmt.Println(children[1])
        new_pop.individuals = append(new_pop.individuals, children[0])
        new_pop.individuals = append(new_pop.individuals, children[1])
        new_pop.size += 2
    }
    new_pop.individuals = append(new_pop.individuals, createGenome(new_pop.individuals[0].genome_len))
    new_pop.individuals = append(new_pop.individuals, createGenome(new_pop.individuals[0].genome_len))
    new_pop.size += 2

    return &new_pop
}

func cross(parent1 *Genome, parent2 *Genome) [2]*Genome {
    var children [2]*Genome
    //fmt.Println(parent1.genome_len)
    var rand_cross int = rand.Intn(parent1.genome_len)

    //fmt.Println(rand_cross)
    var child1_genome []int
    var child2_genome []int
    child1_genome = append(child1_genome, parent1.genome[:rand_cross]...)
    child1_genome = append(child1_genome, parent2.genome[rand_cross:]...)
    if rand.Intn(100) < 1 {
        child1_genome[rand.Intn(len(child1_genome))] = rand.Intn(512) - 256
    }

    child2_genome = append(child2_genome, parent2.genome[:rand_cross]...)
    child2_genome = append(child2_genome, parent1.genome[rand_cross:]...)
    if rand.Intn(100) < 1 {
        child2_genome[rand.Intn(len(child2_genome))] = rand.Intn(512) - 256
    }
    children[0] = createGenomeFromArray(child1_genome)
    children[1] = createGenomeFromArray(child2_genome)
    return children
}
