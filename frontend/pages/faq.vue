<template>
  <div class="bg">
    <div class="faq">
      <h1 class="faq-title">Frequently Asked Questions</h1>
      <div class="search-form-wrapper">
        <v-text-field
          placeholder="Searching for .."
          class="search-form"
          v-model="filter"
        >

        </v-text-field>
      </div>
      <div class="content">

        <template v-for="(subsection, s) in filteredItems">
          <div :key="s">
            <h2 class="sub-section">{{(s + 1) + ' - ' + subsection.title}}</h2> 
            <v-divider></v-divider>
            <template v-for="(question, i) in subsection.items">
              <div :key="i">
                <Drawer :selected="selectedHash" :id="(s + 1) + '.' +(i + 1) + ' - '" :title="question.title" :content="question.content" :references="question.references" />
              </div>
            </template>
          </div>
        </template>

      </div>
    </div>
  </div>
</template>

<script>
import Drawer from "../components/Drawer.vue"

export default {
  name: "faq",
  layout: "faq-layout",
  props: [],
  components: {
    Drawer
  },
  data() {
    return {
      selectedHash: "",
      items: [
              {
                "title": "Core Protocol",
                "items": [
                  {
                    "title": "What are zk-SNARKs?",
                    "content": "Zk-SNARK is an acronym that stands for “Zero-Knowledge Succinct Non-Interactive Argument of Knowledge.” A zk-SNARK is a cryptographic proof that allows one party to prove it possesses certain information without revealing that information. This proof is made possible using a secret key created before the transaction takes place.",
                    "references": [
                      "https://minaprotocol.com/blog/what-are-zk-snarks",
                      "https://www.youtube.com/watch?v=GvwYJDzzI-g",
                      "https://vitalik.ca/general/2021/01/26/snarks.html"
                    ]
                  },
                  {
                    "title": "What is Minas TPS?",
                    "content": "Mina’s current throughput limit is about 1 tps. This was an intentional choice, independent of what the architecture is actually capable of. When deciding what to work on, we’ve chosen to prioritize use cases and decentralization features over a high throughput. Please refer to the blog post from Evan below.",
                    "references": [
                      "https://medium.com/minaprotocol/answering-community-questions-and-whats-ahead-for-mina-d771fa94489b",
                    ]
                  },
                  {
                    "title": "Where can I learn more about the tech behind Mina?",
                    "content": "Mina Protocol has an official YouTube series explaining the tech behind the protocol. Check out the links below.",
                    "references": [
                      "https://www.youtube.com/c/MinaProtocol/videos",
                      "https://minaprotocol.com/get-started",
                      "https://minaprotocol.com/tech",
                      "https://docs.minaprotocol.com/en"
                    ]
                  },
                  {
                    "title": "What makes Mina special?",
                    "content": "Mina is the first layer one protocol that utilizes zk-SNARKs. Zk-SNARKs allow for high scalability, security and decentralisation - trying to solve the 'Blockchain Trilemma'.",
                    "references": [
                      "https://minaprotocol.com/blog/solving-the-scalability-trilemma",
                      "https://vitalik.ca/general/2021/01/26/snarks.html"
                    ]
                  },
                  {
                    "title": "Why are Archive Nodes stored on Google Cloud?",
                    "content": "They aren't. Everyone can run an archive node and many do, they aren't stored on Google Cloud or any similar centralized cloud service.",
                    "references": ["https://medium.com/minaprotocol/answering-community-questions-and-whats-ahead-for-mina-d771fa94489b"]
                  },
                  {
                    "title": "Are Archive Node essential for the network to run?",
                    "content": "No, they are not essential for the network to work as intended. Archive nodes are simply a node that provides convienient benefits of collecting the transaction history for services like Wallets or Blockexplorers.",
                    "references": ["https://docs.minaprotocol.com/en/advanced/archive-node"]
                  },
                  {
                    "title": "Where can I find a roadmap?",
                    "content": "Check out the Product Priorities blog post below.",
                    "references": [
                      "https://minaprotocol.com/blog/mina-protocol-product-priorities-mina-foundation-mission"
                    ]
                  }
                ]
              },
              {
                "title": "Tokenomics and Staking",
                "items": [
                  {
                    "title": "How can I stake my Mina?",
                    "content": "All available wallets for the Mina Protocol allow for easy staking by pressing a button and selecting your staking pool of choice. Please refer to the guide below.",
                    "references": ["https://www.aurowallet.com/mina-staking-guide/"]
                  },
                  {
                    "title": "The Wallet creation fee is too high - can it be changed?",
                    "content": "Nothing is set in stone - it requires just a small fork. And as long as the changes are as minor and undisputed as changing the fee there shouldn't be an issue with forking for that reason. Time will tell, but it can be changed.",
                    "references": []
                  },
                  {
                    "title": "Why is the circulating supply increasing so rapidly? Is someone dumping coins?",
                    "content": "Please refer to the vesting and inflation schedule.",
                    "references": [
                      "https://minaprotocol.com/blog/mina-token-distribution-and-supply"
                    ]
                  },
                  {
                    "title": "Will Mina start burning coins?",
                    "content": "Maybe. Check out the interview with Evan below.",
                    "references": ["https://www.youtube.com/watch?v=x3sLDXHAU-s&t=3136s"]
                  },
                  {
                    "title": "Why is the inflation rate so high?",
                    "content": "A high inflation rate helps to increase the stake on the network and therefor secures it. High inflation incentivises holders to stake their coins.",
                    "references": [
                      "https://minaprotocol.com/blog/mina-token-distribution-and-supply",
                      "https://minaprotocol.com/wp-content/uploads/economicsWhitepaper.pdf"
                    ]
                  },
                  {
                    "title": "When will I start receiving rewards?",
                    "content": "After two full epochs.",
                    "references": []
                  },
                  {
                    "title": "For how long are my coins locked when staking?",
                    "content": "Coins aren't locked, you only delegate your rights to a staking pool. You can access coins at any time.",
                    "references": []
                  },
                  {
                    "title": "Wen moon?",
                    "content": "No.",
                    "references": ["https://www.youtube.com/watch?v=dQw4w9WgXcQ"]
                  }
                ]
              },
              {
                "title": "Wallets",
                "items": [
                  {
                    "title": "Can I use MetaMask?",
                    "content": "No, Mina is not an BEP20 or ETH20 token - it is required to use a Wallet specifically for the Mina Protocol.",
                    "references": ["https://www.aurowallet.com/", "https://clor.io/"]
                  },
                  {
                    "title": "Will there be an official Wallet by the Foundation?",
                    "content": "Probably not. Community wallets (Auro and Clorio) have been audited and are secure according to the audit. However, Mina has an official ClientSDK and a CLI where you can generate a key yourself.",
                    "references": ["https://www.aurowallet.com/", "https://clor.io/"]
                  },
                  {
                    "title": "What wallet should I use?",
                    "content": "Both Auro and Clorio wallet have been audited. Alternatively you could also use the Mina ClientSDK or CLI.",
                    "references": ["https://www.aurowallet.com/", "https://clor.io/"]
                  }
                ]
              },
              {
                "title": "DeFi and NFTs",
                "items": [
                  {
                    "title": "Will NFTs and DeFi be possible on Mina?",
                    "content": "Minas version of Smart Contracts, Snapps, are still being worked on. Eventually DeFi and dApps will also come to Mina. Please refer to the Snapps section.",
                    "references": []
                  }
                ]
              },
              {
                "title": "Snapps and SnarkyJS",
                "items": [
                  {
                    "title": "What are Snapps?",
                    "content": "Snapps are Minas version of Smart Contracts with zk-SNARKs.",
                    "references": ["https://www.youtube.com/watch?v=H_JQjPDwAH0"]
                  },
                  {
                    "title": "Can I use Solidity to write Snapps?",
                    "content": "No, Snapps will be written in TypeScript.",
                    "references": ["https://github.com/o1-labs/snarkyjs"]
                  },
                  {
                    "title": "What is SnarkyJS?",
                    "content": "A TypeScript framework for zk-SNARKs and Snapps.",
                    "references": ["https://github.com/o1-labs/snarkyjs"]
                  }
                ]
              }
            ],
      filter: "",
    }
  },
  async created() {
    // const response = await fetch('/faq.json');
    // let json = await response.json()
    // this.items = json
    this.selectedHash = (this.$nuxt.$route.hash).substring(1);
  },
  computed: {
    filteredItems() {
      if(this.filter == "") {
        return this.items
      }
      
      // let filteredSections = new Array()

      // this.items.forEach(section => {
      //   if (section.title.toLowerCase().includes(this.filter.toLowerCase()))
      //     filteredSections.push(section)
      // })


      // return filteredSections


      return this.items.filter(item => item.items.some(s => s.title.toLowerCase().includes(this.filter.toLowerCase())));

    }  
  }
}
</script>


<style scoped>
.bg {
  background-color: rgba(0, 0, 0, 0.186);
  border-radius: 5px;
  height: 100%;
  display: flex;
  padding: 15px;
}
.faq {
  justify-content: center;
  text-align: center;
  width: 100%;
  height: 100% !important;
}

.content {
  width: 100%;
  justify-content: left;
  text-align: left;
  float: left;
  margin-bottom: 25px;
}

.sub-section {
  margin-top: 15px;
  margin-left: 15px;
  font-weight: 400;
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}

.faq-title {
  font-size: 1.5rem;
  font-weight: 400;
  color: rgb(255, 255, 255);
  -webkit-text-stroke: 0.5px rgb(0, 0, 0);
}

.search-form-wrapper {
  justify-content: center;
  align-items: center;
}

.search-form {
  width: 400px;

}

</style>