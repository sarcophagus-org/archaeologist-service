apt-get update
apt-get install build-essential
sudo snap install core; sudo snap refresh core
rm -rf /root/go && snap install go --classic
export PATH=$PATH:/root/go/bin
git clone https://github.com/decent-labs/airfoil-sarcophagus-archaeologist-service /home/archservice
cp /home/archservice/config.example.yml /home/archservice/config.yml
cp /home/archservice/config.yml /root/go/bin/config.yml
cd /home/archservice && go install

sudo apt-get install nginx
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
sudo certbot certonly --nginx
nano /etc/nginx/sites-available/reverse-proxy.conf
service nginx reload

nano /etc/systemd/system/archservice.service
