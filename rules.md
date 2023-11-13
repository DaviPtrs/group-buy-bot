# Group Buy do S&M

Com a proeza do estado brasileiro de taxar compras do aliexpress direto na fonte, precisamos de outra forma de comprar coisas de qualidade por um preço ok. O pulo do gato do aliexpress pra pagar menos imposto (ou até nenhum) era sub declarar o valor dos produtos, pra quando chegar na alfandega o fiscal achar que se trata de algo mixaria e, caso resolvesse taxar mesmo assim, voce seria taxado em cima do valor declarado (que era quase sempre bem baixo).

Agora que o **roubo** (ou *imposto*, em estatês) é direto no carrinho, pro aliexpress isso não funciona mais. Porém, compras dos EUA em que se usa um redirecionador (ou até uma pessoa física que te envie de lá) sempre foi possível sub declarar, dado que quem preenche a declaração é quem estiver usando o serviço.

Pra quem não sabe, um redirecionador é geralmente uma empresa que te fornece um endereço no país (no nosso caso, nos EUA) aonde você pode comprar em lojas americanas diferentes e enviar pra lá. Quando sua encomenda chega nesse endereço, eles registram seus produtos numa plataforma aonde você pode pedir pra jogar tudo numa caixa só e enviar pra qualquer outro país. Nesse envio, você escolhe o frete e também preenche a declaração, dessa forma você pode sub declarar a vontade pra pagar menos imposto (tendo bom senso e malícia, não adianta declarar um iphone de 1000 dolares como "gift $2" e achar que vai passar).

A questão do redirecionamento não ser tão popular quanto o aliexpress foi é que o frete dos EUA para o Brasil é feito via transportadora aérea tercerizada (ao contrário do aliexpress, que mandava avião de carga pro Brasil algumas vezes por semana), isso faz com que o valor do frete seja mais caro.

Portanto, viemos com a ideia do group buy justamente pra conseguir baratear esse custo de comprar dos EUA, juntando várias pessoas e rachando o frete entre elas. Com 2 pessoas ou mais, já vale a pena fazer um envio.

Depois de terminar a leitura de todas essas mensagens, reaja a essa mensagem com um :thumbs-up: pra confirmar que leu tudo.


## Como funciona

Esse bot pode ser resumido em uma lista de desejos coletiva, aonde quando atingindo 2 pessoas diferentes ou mais que queiram comprar coisas de fora e, mediante o OK de todas as pessoas envolvidas, o admin do Group Buy (eu, no caso) organizará o envio desses produtos pro Brasil, fazendo a racha dos custos com todos os envolvidos. (Inclusive esse bot poderia ser só um Google Form, mas ai qual seria a graça?)

Pra usar o bot você:

1. Nesse canal, digite `/add`
2. Aparecerá um formulário pra você inserir o link e informações do produto que você quer comprar
3. Seu requerimento irá para a análise (obviamente nem tudo vale a pena importar por esse método, cheque a seção "**Regras**")
4. Se for negado, o BOT te enviará uma mensagem explicando o motivo.
5. Se for aprovado, seu produto irá pra uma lista de desejo.
6. Assim que a lista atingir 2 pessoas diferentes com items nela, o bot enviará uma mensagem pra todos os envolvidos.
7. Os envolvidos podem combinar quando fazer a compra dos produtos ou, esperar pra entrar mais uma pessoa (assim o frete fica ainda mais barato)
8. Pra cada produto que chegar na redirecionadora o admin (eu) receberá um email, mas por precaução, quando seu produto chegar no destino americano (a responsabilidade de acompanhar o rastreio é de quem comprou) avise o admin no privado.
9. Assim que os envolvidos decidirem fazer o envio pro Brasil, o admin vai calcular os custos de importação e passar o valor que cada um deve pagar.
10. Caso o comprador more fora do estado de onde a caixa chegará (cheque a seção "**Destino**"), a pessoa que receber a encomenda irá fazer o envio pra cada comprador. Não preciso nem falar que o custo desse frete também é por conta do comprador.

## Envio

O valor do envio dos EUA pro Brasil é por peso, e a unidade de medida americana pra peso, ao invés de gramas e kilos, é libras (ou pounds). É interessante que você saiba quanto pesa o item que você quer comprar, assim você pode ter uma noção de quanto pagará de frete. Lembrando que o custo total será dividido proporcionalmente, isso é, quem ta importando objeto mais leve vai pagar menos.

-   16 Ounces = 1 Pound/libra
-   1 Kg = 2,205 libras
-   1 Libra = 453,59 gramas

Pra fazer uma simulação do custo de envio, use [essa calculadora](https://zip4me.com/simulator-freight.html)

Selecione:
- Brazil (duh)
- Tipo de entrega "Normal"
- O envio é o "ZipStandard" (é o mais barato)

## Endereço

Esse é o endereço que você deverá enviar seu produto **(atenção ao Addr 2, ou complemento, é por ele que sabemos que a caixa vai pro meu nome)**:

**Nome:** Davi Petris

**Addr 1:** 13151 NE Airport Way Bldg 14 

**Addr 2:** MCC 130565

**City:** Portland

**State/Province:** Oregon - OR

**Zip code:** 97230-1036

**Country:** US


##  Regras e Termos

-   Algumas lojas e marcas além de vender no aliexpress também possuem lojas próprias. Muitas dessas lojas, por serem chinesas, enviam direto na China quando você compra na loja deles.
    
    Verifique se é o caso do produto que você quer comprar. Também é válido enviar um email pro "Contact Us" da loja perguntando se eles declaram um valor pequeno na hora de enviar. 
    
    Muitas lojas já declaram um valor menor sem que você peça, outras tem um campo específico pra você preencher o valor a ser declarado, outras se recusam a fazer essa prática. Então é sempre bom confirmar.
-   Pelo custo do frete ser calculado com base no peso, pode não valer a pena importar seu produto se ele for muito pesado ou muito grande.
-   Produtos que chamam muita atenção (ex: coisas da Apple, notebooks, tablets, celulares) podem ser barrados do Group buy, justamente por acabar comprometendo a caixa inteira.
    
    Se um fiscal ver um iPhone sub declarado, a caixa inteira pode rodar ou pagar multa por conta disso)
-   Como sabem, os **parasitas** (ou *receita federal*, em estatês) tem critérios completamente subjetivos pra analisar uma caixa. 
    
    Eles podem não taxar, podem taxar em cima do valor declarado, podem riscar o valor (por achar que o valor está errado) e taxar em cima do valor novo, podem abrir a caixa pra conferir o conteúdo, podem multar (além de taxar) e podem inclusive barrar a entrada da caixa no Brasil (nesse caso ela volta pra redirecionadora mas isso é bem demorado). 
    
    Portanto, por favor, saiba que existem riscos. Ninguém vai arcar com custo nenhum sozinho.
-   Caso a caixa não seja taxada, o valor relacionado ao pagamento do imposto será devolvido.
-   A redirecionadora dá um prazo de 90 dias para o envio dos items ao Brasil, caso passe desse prazo, os produtos são descartados. 
    
    A não ser tenha algum maluco pra pagar 1$ pra cada dia a mais de armazenamento.
-   A caixa **NÃO SERÁ ENVIADA** ao Brasil até que todos paguem os custos devidos.
-   No momento do OK para envio ao Brasil, a partir do recebimento dos valores a serem pagos, os envolvidos terão o prazo de 5 dias pra pagar. 
    
    Caso a lista tenha mais de 2 pessoas, quem não pagar dentro do prazo ficará de fora do envio, podendo até perder o item caso não envie dentro do prazo de armazenamento.
-   Caso alguem haja de má fé, será banido de participar do Group Buy ou até mesmo banido do servidor (fica a critério do dono)
-   Não é possível declarar um produto pra pagar 0% de imposto, por favor não insista. A caixa ser taxada ou não vai depender de sorte e, ultimamente, todas as caixas vindas dos EUA estão sendo taxadas.
-   A princípio, a responsabilidade pelo ato de comprar o produto é de quem pediu, porém caso você tenha dificuldades (costumava comprar no aliexpress via boleto ou pix), nos contate que te ajudamos nisso (**não abuse da minha fucking boa vontade**)

## Destino

O destino da caixa será sempre Vitória - Espírito Santo. Pois é aonde o organizador do Group Buy mora (e eu não to recebendo nada por isso então nada mais justo), além do Espirito Santo não cobrar ICMS sobre produtos importados, e portanto, o imposto será menor.

Esse ponto tem uma exceção: Caso não tenham produtos meus no Group Buy (vai ser raro mas pode acontecer) e algum dos compradores morar em um estado que também não cobre ICMS, a caixa será enviada pra casa desse comprador.