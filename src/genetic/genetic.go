package genetic
import "math/rand"
import "time"
import "fmt"

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
        g.genome = append(g.genome, rand.Intn(255))
    }
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

func PrintGenome(genome *Genome) {
    fmt.Println(genome.genome, genome.fitness)
}

func GetIndividuals(pop *Population) []*Genome{
    return pop.individuals
}

func (genome *Genome) SetFitness(value int) {
    genome.fitness = value
}
