#include <iostream>
#include <vector>
#include <list>
// 400 p, 71864m
using namespace std;
using ll = long long;

int main() {
    int P = 416;
    int M = 71671;

    vector<ll> score(P, 0);
    list<ll> A;
    A.push_front(0);
    auto p = A.begin();
    for(int m = 1; m <= M; m++){
       if(m%23 != 0){
           for(int i = 0; i < 2; i++){
               p++;
               if(p == A.end()){
                   p = A.begin();
               }
            }
        A.insert(p, m);
        p--;

       }
       else {
           score[m%P] += m;
           for(int i = 0; i < 7; i++){
               if(p == A.begin()){
                    p = A.end();
               }
               p--;
            }
           score[m%P] += *p;
           A.erase(p);
           p++;
           if(p == A.end()) { p = A.begin(); }
       }    
    }

    ll maxscore = 0;

    for(ll i=0; i<P; i++){
       if(score[i] > maxscore){ maxscore = score[i]; } 
    }
    cout << maxscore << endl;
}
